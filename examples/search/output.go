package main

import (
	"context"
	"fmt"
)

// printHighlightedLine prints a line with ANSI color highlighting.
// In a production version, you'd re-apply the regex pattern to highlight
// only the matching portions, but this simplified version highlights the entire line.
func printHighlightedLine(lineNum int, line string) (err error) {
	var highlighted string

	// Apply ANSI color codes for highlighting
	highlighted = highlightMatch(line)
	fmt.Printf("%d:  %s\n", lineNum, highlighted)

	return err
}

// highlightMatch applies ANSI color codes to highlight text.
// Uses red color (code 31) with reset (code 0) afterward.
// In a more sophisticated version, this would re-apply the regex pattern
// to highlight only the matching portions within the line.
func highlightMatch(line string) (result string) {
	// Simple highlighting: make entire line red
	// \033[31m = red text, \033[0m = reset to normal
	result = fmt.Sprintf("\033[31m%s\033[0m", line)
	return result
}

// printMatch formats and prints a single match result.
// It shows the file path, context lines, and highlights the matching line.
// The format mimics grep's output style for familiarity.
func printMatch(match Match) (err error) {
	// Print file path header
	fmt.Printf("\n%s:\n", match.FilePath)

	// Print line before match (if exists)
	if match.Before != "" {
		fmt.Printf("%d-  %s\n", match.LineNumber-1, match.Before)
	}

	// Print the matching line with highlighting
	err = printHighlightedLine(match.LineNumber, match.Line)
	if err != nil {
		goto end
	}

	// Print line after match (if exists)
	if match.After != "" {
		fmt.Printf("%d+  %s\n", match.LineNumber+1, match.After)
	}

end:
	return err
}

// outputHandler is responsible for receiving match results and printing them.
// It runs in its own goroutine and provides a centralized place for output
// formatting, which prevents garbled output from concurrent goroutines.
//
// This demonstrates the fan-in pattern where multiple producers send to
// a single consumer for coordinated processing.
func outputHandler(ctx context.Context, matchChan <-chan Match) (err error) {
	if verbose {
		fmt.Printf("[TRACE] Output handler started\n")
	}

	// Run until channel is closed or context is cancelled
	for {
		select {
		case match, ok := <-matchChan:
			// Channel closed, no more matches
			if !ok {
				if verbose {
					fmt.Printf("[TRACE] Match channel closed, output handler exiting\n")
				}
				goto end
			}
			if verbose {
				fmt.Printf("[TRACE] Received match from %s:%d\n", match.FilePath, match.LineNumber)
			}
			// Print the match result
			err = printMatch(match)
			if err != nil {
				if verbose {
					fmt.Printf("[TRACE] Error printing match: %v\n", err)
				}
				goto end
			}
		case <-ctx.Done():
			// Context cancelled, stop processing
			if verbose {
				fmt.Printf("[TRACE] Output handler context cancelled\n")
			}
			err = ctx.Err()
			goto end
		}
	}

end:
	if verbose {
		fmt.Printf("[TRACE] Output handler finished\n")
	}
	return err
}
