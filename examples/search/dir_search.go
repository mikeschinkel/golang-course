package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/sync/errgroup"
)

// DirSearch encapsulates all the state and configuration needed for a directory search.
// This eliminates prop drilling and provides a clean, testable interface.
type DirSearch struct {
	searchDir     string
	glob          string
	pattern       *regexp.Regexp
	workerLimiter chan struct{}
	matchChan     chan Match
	verbose       bool
}

// NewDirSearch creates a new directory search instance with the specified configuration.
func NewDirSearch(searchDir, glob string, pattern *regexp.Regexp, maxWorkers int, verbose bool) *DirSearch {
	return &DirSearch{
		searchDir:     searchDir,
		glob:          glob,
		pattern:       pattern,
		workerLimiter: make(chan struct{}, maxWorkers),
		matchChan:     make(chan Match, maxWorkers),
		verbose:       verbose,
	}
}

// Run executes the directory search and returns any error encountered.
// It coordinates the overall search operation by setting up:
// 1. Context and cancellation handling
// 2. Channel for collecting results
// 3. Worker limiting semaphore
// 4. Two main goroutines: one for output, and one entry-point for searching that spawns other goroutines
func (ds *DirSearch) Run(ctx context.Context) (err error) {
	var cancel context.CancelFunc
	var g *errgroup.Group

	if ds.verbose {
		fmt.Printf("[TRACE] Starting search in %s with pattern %s\n", ds.searchDir, ds.pattern.String())
	}

	// Create cancellable context for coordinating shutdown
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	// Create errgroup for coordinated goroutine management
	g, ctx = errgroup.WithContext(ctx)

	if ds.verbose {
		fmt.Printf("[TRACE] Starting output handler goroutine\n")
	}

	// Start the output handler goroutine
	g.Go(func() error {
		return ds.outputHandler(ctx)
	})

	if ds.verbose {
		fmt.Printf("[TRACE] Starting search goroutine\n")
	}

	// Start the search goroutine
	g.Go(func() error {
		defer func() {
			if ds.verbose {
				fmt.Printf("[TRACE] Closing match channel\n")
			}
			close(ds.matchChan)
		}()
		return ds.searchDirectory(ctx, ds.searchDir)
	})

	if ds.verbose {
		fmt.Printf("[TRACE] Waiting for all goroutines to complete\n")
	}

	err = g.Wait()

	if ds.verbose {
		fmt.Printf("[TRACE] All goroutines completed\n")
	}

	return err
}

// searchDirectory recursively searches a single directory level.
// It spawns goroutines for:
// 1. Each subdirectory (recursive search)
// 2. Each file that matches the glob pattern
//
// This function demonstrates the concurrent directory traversal pattern
// where each directory level manages its own set of worker goroutines.
func (ds *DirSearch) searchDirectory(ctx context.Context, dir string) (err error) {
	var g *errgroup.Group
	var entries []os.DirEntry

	if ds.verbose {
		fmt.Printf("[TRACE] Entering directory: %s\n", dir)
	}

	// Check for cancellation before starting expensive directory operations
	select {
	case <-ctx.Done():
		if ds.verbose {
			fmt.Printf("[TRACE] Context cancelled in directory: %s\n", dir)
		}
		err = ctx.Err()
		return
	default:
	}

	// Create errgroup for this directory level's goroutines
	// The context from errgroup will be cancelled if any child goroutine fails
	g, ctx = errgroup.WithContext(ctx)

	// Read directory contents - this can be slow for large directories
	entries, err = os.ReadDir(dir)
	if err != nil {
		if ds.verbose {
			fmt.Printf("[TRACE] Cannot read directory %s: %v\n", dir, err)
		}
		// Don't fail the entire search for one unreadable directory
		// This handles permission errors, broken symlinks, etc.
		return nil
	}

	if ds.verbose {
		fmt.Printf("[TRACE] Found %d entries in %s\n", len(entries), dir)
	}

	// Process each directory entry (files and subdirectories)
	err = ds.processDirectoryEntries(ctx, g, dir, entries)
	if err != nil {
		if ds.verbose {
			fmt.Printf("[TRACE] Error processing entries in %s: %v\n", dir, err)
		}
		goto end
	}

	if ds.verbose {
		fmt.Printf("[TRACE] Waiting for goroutines in %s\n", dir)
	}

	// Wait for all child goroutines (subdirectories and files) to complete
	// This ensures we don't return until all work for this directory is done
	err = g.Wait()

	if ds.verbose {
		fmt.Printf("[TRACE] Finished directory: %s\n", dir)
	}

end:
	return err
}

// processDirectoryEntries iterates through directory entries and spawns goroutines
// for subdirectories and matching files. This is where the core concurrency happens.
//
// Key concurrency concepts demonstrated:
// 1. Goroutine closure variable capture (must capture loop variables by value)
// 2. Worker limiting using channel-based semaphores
// 3. Selective processing (skip certain directories, match files by glob)
func (ds *DirSearch) processDirectoryEntries(ctx context.Context, g *errgroup.Group, dir string, entries []os.DirEntry) (err error) {
	var fullPath string
	var matched bool

	// Iterate through each entry in the directory
	for _, entry := range entries {
		// Check for cancellation on each iteration
		// This allows quick response to Ctrl-C even in large directories
		select {
		case <-ctx.Done():
			err = ctx.Err()
			goto end
		default:
		}

		// Build full path for this entry
		fullPath = filepath.Join(dir, entry.Name())

		// Handle directories: recurse into subdirectories
		if entry.IsDir() {
			// Skip directories we don't want to search (optimization)
			if shouldSkipDirectory(entry.Name()) {
				continue
			}

			// Spawn goroutine for recursive directory search
			// CRITICAL: Must capture fullPath by value to avoid closure bug
			// Without this pattern, all goroutines would search the same directory
			capturedPath := fullPath // Capture by value
			g.Go(func() error {
				// Recursively search the subdirectory - no worker limiting for directory traversal
				return ds.searchDirectory(ctx, capturedPath)
			})
			continue
		}

		// Handle files: check if they match the glob pattern
		matched, err = filepath.Match(ds.glob, entry.Name())
		if err != nil {
			goto end
		}

		// Skip files that don't match the pattern
		if !matched {
			continue
		}

		// Spawn goroutine to search this file
		// Same closure pattern as directories to avoid variable capture bug
		capturedPath := fullPath // Capture by value
		g.Go(func() error {
			// Acquire worker slot inside the goroutine
			ds.workerLimiter <- struct{}{}
			defer func() {
				<-ds.workerLimiter
			}()
			// Search the file for matches
			return ds.searchFile(ctx, capturedPath)
		})
	}

end:
	return err
}

// searchFile searches a single file for pattern matches.
// This function demonstrates several important patterns:
// 1. Resource validation (file type, size, permissions)
// 2. Binary file detection
// 3. Streaming file processing with context cancellation
// 4. Error handling without failing the entire search
func (ds *DirSearch) searchFile(ctx context.Context, filePath string) (err error) {
	var file *os.File
	var scanner *bufio.Scanner
	var lines []string
	var lineNum int
	var stat os.FileInfo
	var isTextFile bool
	var buf []byte

	if ds.verbose {
		fmt.Printf("[TRACE] Searching file: %s\n", filePath)
	}

	// Check for cancellation before starting expensive file operations
	select {
	case <-ctx.Done():
		if ds.verbose {
			fmt.Printf("[TRACE] Context cancelled for file: %s\n", filePath)
		}
		err = ctx.Err()
		goto end
	default:
	}

	// Get file information without opening it
	// This handles broken symlinks gracefully
	stat, err = os.Stat(filePath)
	if err != nil {
		if ds.verbose {
			fmt.Printf("[TRACE] Cannot stat file %s: %v\n", filePath, err)
		}
		// Skip files we can't stat (broken symlinks, permission issues, etc.)
		// Don't fail the entire search for one problematic file
		goto end
	}

	// Double-check that it's actually a file (race condition protection)
	if stat.IsDir() {
		if ds.verbose {
			fmt.Printf("[TRACE] Skipping directory misidentified as file: %s\n", filePath)
		}
		goto end
	}

	if stat.Size() > 50*1024*1024 {
		if ds.verbose {
			fmt.Printf("[TRACE] Skipping large file: %s (%d bytes)\n", filePath, stat.Size())
		}
		goto end
	}

	// Open the file for reading
	file, err = os.Open(filePath)
	if err != nil {
		if ds.verbose {
			fmt.Printf("[TRACE] Cannot open file %s: %v\n", filePath, err)
		}
		// Skip files we can't open (permissions, broken symlinks, etc.)
		goto end
	}
	// Ensure file is closed, even if errors occur
	// Clear Path style: handle close errors properly
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// Check if file appears to be text (binary file detection)
	isTextFile, err = isLikelyTextFile(file)
	if err != nil || !isTextFile {
		if ds.verbose {
			if err != nil {
				fmt.Printf("[TRACE] Error checking if %s is text: %v\n", filePath, err)
			} else {
				fmt.Printf("[TRACE] Skipping binary file: %s\n", filePath)
			}
		}
		goto end
	}

	// Reset file position after checking (seek back to beginning)
	_, err = file.Seek(0, 0)
	if err != nil {
		if ds.verbose {
			fmt.Printf("[TRACE] Cannot seek file %s: %v\n", filePath, err)
		}
		goto end
	}

	// Set up scanner with increased buffer for long lines
	scanner = bufio.NewScanner(file)
	buf = make([]byte, 0, 64*1024) // Initial buffer size
	scanner.Buffer(buf, 1024*1024) // Max token size (1MB)

	// Track lines for context (before/after match)
	lines = make([]string, 0)
	lineNum = 0

	// Scan file line by line
	for scanner.Scan() {
		// Check for cancellation periodically during file processing
		// This allows responsive cancellation even for large files
		select {
		case <-ctx.Done():
			if ds.verbose {
				fmt.Printf("[TRACE] Context cancelled while scanning file: %s\n", filePath)
			}
			err = ctx.Err()
			goto end
		default:
		}

		lineNum++
		line := scanner.Text()
		lines = append(lines, line)

		// Check if current line matches the pattern
		if !ds.pattern.MatchString(line) {
			continue
		}

		if ds.verbose {
			fmt.Printf("[TRACE] Found match in %s at line %d\n", filePath, lineNum)
		}

		// Send match result to output handler
		// This demonstrates channel communication between goroutines
		err = ds.sendMatch(ctx, filePath, lineNum, lines, len(lines)-1)
		if err != nil {
			if ds.verbose {
				fmt.Printf("[TRACE] Error sending match for %s: %v\n", filePath, err)
			}
			goto end
		}
	}

	// Check for scanner errors (EOF is normal, others are problems)
	err = scanner.Err()
	if err != nil {
		if ds.verbose {
			fmt.Printf("[TRACE] Scanner error for %s: %v\n", filePath, err)
		}
	} else if ds.verbose {
		fmt.Printf("[TRACE] Finished searching file: %s\n", filePath)
	}

end:
	return err
}

// sendMatch creates a Match struct and sends it to the output handler.
// It includes context (before/after lines) and handles channel communication
// with proper cancellation support.
func (ds *DirSearch) sendMatch(ctx context.Context, filePath string, lineNum int, lines []string, matchIndex int) (err error) {

	var match Match
	var before string
	var after string

	// Get context lines (before and after the match)
	if matchIndex > 0 {
		before = lines[matchIndex-1]
	}

	// Check if there's a line after (avoid index out of bounds)
	if matchIndex < len(lines)-1 {
		after = lines[matchIndex+1]
	}

	// Create match result
	match = Match{
		FilePath:   filePath,
		LineNumber: lineNum,
		Line:       lines[matchIndex],
		Before:     before,
		After:      after,
		IsMatch:    true,
	}

	// Send to output handler with cancellation support
	// This demonstrates non-blocking channel communication
	select {
	case ds.matchChan <- match:
		// Successfully sent
	case <-ctx.Done():
		// Context cancelled, abort sending
		err = ctx.Err()
		goto end
	}

end:
	return err
}

// outputHandler receives match results and prints them.
// Accesses matchChan via receiver.
func (ds *DirSearch) outputHandler(ctx context.Context) (err error) {
	if ds.verbose {
		fmt.Printf("[TRACE] Output handler started\n")
	}

	for {
		select {
		case match, ok := <-ds.matchChan:
			if !ok {
				if ds.verbose {
					fmt.Printf("[TRACE] Match channel closed, output handler exiting\n")
				}
				goto end
			}
			if ds.verbose {
				fmt.Printf("[TRACE] Received match from %s:%d\n", match.FilePath, match.LineNumber)
			}
			err = printMatch(match)
			if err != nil {
				if ds.verbose {
					fmt.Printf("[TRACE] Error printing match: %v\n", err)
				}
				goto end
			}
		case <-ctx.Done():
			if ds.verbose {
				fmt.Printf("[TRACE] Output handler context cancelled\n")
			}
			err = ctx.Err()
			goto end
		}
	}

end:
	if ds.verbose {
		fmt.Printf("[TRACE] Output handler finished\n")
	}
	return err
}
