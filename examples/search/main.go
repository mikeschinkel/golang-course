// Package main implements a concurrent grep-like tool that searches for regex patterns
// across files in a directory tree using goroutines for parallel processing.
//
// Key features:
// - Concurrent directory traversal (one goroutine per directory)
// - Concurrent file searching (one goroutine per matching file)
// - Worker limiting to prevent resource exhaustion
// - Context-based cancellation for clean Ctrl-C handling
// - Channel-based result coordination
// - Binary file detection and skipping
// - Symlink handling
//
// This implementation demonstrates advanced Go concurrency patterns including:
// - errgroup for coordinated goroutine management
// - Context cancellation propagation
// - Channel-based semaphores for resource limiting
// - Proper closure variable capture in goroutines
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

// maxWorkers defines the maximum number of concurrent goroutines allowed.
// This prevents resource exhaustion on large directory trees.
const maxWorkers = 100

// verbose controls whether detailed tracing output is enabled.
// Set by the -v command line flag.
var verbose bool

// Match represents a single search result containing the matched line
// and its surrounding context (one line before and after).
type Match struct {
	FilePath   string // Full path to the file containing the match
	LineNumber int    // Line number where the match was found (1-based)
	Line       string // The actual line containing the match
	Before     string // Line immediately before the match (empty if none)
	After      string // Line immediately after the match (empty if none)
	IsMatch    bool   // Always true for actual matches (used for type safety)
}

// main is the entry point. It follows the Clear Path style with minimal nesting
// and a single error handling path at the end.
func main() {
	var err error
	var pattern *regexp.Regexp
	var searchDir string
	var glob string
	var dirSearch *DirSearch
	var ctx context.Context
	var cancel context.CancelFunc

	// Parse command line arguments and compile the regex pattern
	err = parseArgs(&searchDir, &glob, &pattern)
	if err != nil {
		goto end
	}

	// Create DirSearch instance
	dirSearch = NewDirSearch(searchDir, glob, pattern, maxWorkers, verbose)

	// Set up signal handling
	ctx = context.Background()
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	err = setupSignalHandler(cancel)
	if err != nil {
		goto end
	}

	// Run the search
	err = dirSearch.Run(ctx)

end:
	// Clear Path style error handling: single exit point with proper error reporting
	// Don't treat EOF as an error - it's normal when scanning reaches end of file
	if err != nil && err.Error() != "EOF" {
		// Attempt to write error to stderr, but handle the case where even that fails
		_, err = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err != nil {
			// Fallback to log if stderr write fails
			log.Print(err.Error())
		}
		os.Exit(1)
	}
}

// setupSignalHandler configures graceful shutdown on SIGINT (Ctrl-C) and SIGTERM.
// When a signal is received, it calls the cancel function to trigger context cancellation,
// which propagates through all goroutines for coordinated shutdown.
func setupSignalHandler(cancel context.CancelFunc) (err error) {
	var sigChan chan os.Signal

	// Create buffered channel to receive OS signals
	sigChan = make(chan os.Signal, 1)

	// Register interest in interrupt and termination signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start goroutine to handle signals
	// This runs independently and triggers cancellation when signals arrive
	go func() {
		<-sigChan // Block until signal received
		fmt.Println("Signal received, cancelling...")
		cancel() // Trigger context cancellation
	}()

	return err
}

// parseArgs processes command line arguments and handles the -v verbose flag.
// It supports both directory-only patterns (with trailing slash) and glob patterns.
// Examples:
//
//	search ~/Projects/ "error"           -> search all files in ~/Projects
//	search ~/Projects/*.go "func"        -> search only .go files
//	search -v ~/Projects/ "error"        -> same as first, with verbose output
func parseArgs(searchDir *string, glob *string, pattern **regexp.Regexp) (err error) {
	var pathPattern string
	var dir string
	var file string
	var expandedPath string
	var args []string

	// Filter out the -v flag while preserving argument order
	// This demonstrates slice manipulation and flag processing
	args = make([]string, 0, len(os.Args))
	for _, arg := range os.Args {
		if arg == "-v" {
			verbose = true
			continue // Skip the -v flag, don't add it to filtered args
		}
		args = append(args, arg)
	}

	// Validate we have enough arguments after filtering
	if len(args) < 3 {
		err = fmt.Errorf("usage: %s [-v] <path_pattern> <regex_pattern>", os.Args[0])
		goto end
	}

	pathPattern = args[1]

	// Handle trailing slash as "search everything in this directory"
	// This provides a clean UX: ~/Projects/ means "all files in ~/Projects"
	if strings.HasSuffix(pathPattern, "/") {
		// Remove trailing slash and expand tilde
		expandedPath, err = expandTilde(strings.TrimSuffix(pathPattern, "/"))
		if err != nil {
			goto end
		}
		*searchDir = expandedPath
		*glob = "*" // Match all filesz
	} else {
		// Split path into directory and glob pattern
		// Example: ~/Projects/*.go -> dir="~/Projects", file="*.go"
		dir = filepath.Dir(pathPattern)
		file = filepath.Base(pathPattern)
		expandedPath, err = expandTilde(dir)
		if err != nil {
			goto end
		}
		*searchDir = expandedPath
		*glob = file
	}

	// Compile the regex pattern - this validates it's syntactically correct
	*pattern, err = regexp.Compile(args[2])

end:
	return err
}
