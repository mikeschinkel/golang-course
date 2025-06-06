package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// shouldSkipDirectory determines which directories to skip during traversal.
// This improves performance by avoiding directories that typically contain
// generated files, dependencies, or version control data.
func shouldSkipDirectory(dirName string) (skip bool) {
	switch dirName {
	case ".git": // Git repository data
	case "node_modules": // Node.js dependencies
	case ".svn", ".hg": // Other version control systems
	case "vendor": // Go/PHP dependencies
	case "target": // Rust/Java build output
	case "build", "dist": // Common build directories
	case ".idea", ".vscode": // IDE configuration
		skip = true
	default:
		skip = false
	}
	return skip
}

// expandTilde converts Unix-style ~ home directory notation to full paths.
// Go doesn't expand ~ automatically like shells do, so we handle it manually.
// Supports both "~" (home directory) and "~/path" (path within home directory).
func expandTilde(path string) (result string, err error) {
	var usr *user.User

	// If path doesn't start with ~, return as-is
	if !strings.HasPrefix(path, "~") {
		result = path
		goto end
	}

	// Handle bare ~ (just the home directory)
	if path == "~" {
		usr, err = user.Current()
		if err != nil {
			goto end
		}
		result = usr.HomeDir
		goto end
	}

	// Handle ~/path (path within home directory)
	if strings.HasPrefix(path, "~/") {
		usr, err = user.Current()
		if err != nil {
			goto end
		}
		// Join home directory with the path after ~/
		result = filepath.Join(usr.HomeDir, path[2:])
		goto end
	}

	// If we get here, it's some other ~-prefixed path we don't handle
	result = path

end:
	return result, err
}

// isLikelyTextFile determines if a file is likely to contain text by examining
// the first 512 bytes for null characters. Binary files typically contain many
// null bytes, while text files contain very few or none.
//
// This is a heuristic approach - not 100% accurate but works well in practice.
func isLikelyTextFile(file *os.File) (isText bool, err error) {
	var buf [512]byte
	var n int
	var nullCount int

	// Read first 512 bytes of file
	n, err = file.Read(buf[:])
	if err != nil && n == 0 {
		goto end
	}

	// Reset error if we got some data (EOF after reading is normal)
	if n > 0 {
		err = nil
	}

	// Count null bytes in the sample
	for i := 0; i < n; i++ {
		if buf[i] == 0 {
			nullCount++
		}
	}

	// If more than 1% null bytes, probably binary
	// This threshold works well for most text vs binary classification
	isText = (nullCount * 100 / n) < 1

end:
	return isText, err
}
