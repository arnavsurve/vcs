package cmd

import (
	"bufio"
	"os"
	"path/filepath"
)

// read .vcsignore file and return a list of patterns to ignore
func ReadVCSIgnore() ([]string, error) {
	var patterns []string
	file, err := os.Open(".vcsignore")
	if err != nil {
		return nil, err // return error if file does not exist
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pattern := scanner.Text()
		if pattern != "" {
			patterns = append(patterns, pattern)
		}
	}

	return patterns, scanner.Err()
}

// check if a file path matches patterns to ignore
func IsIgnored(filePath string, ignorePatterns []string) bool {
	for _, pattern := range ignorePatterns {
		matched, err := filepath.Match(pattern, filePath)
		if err != nil {
			continue // skip patterns that cause errors
		}
		if matched {
			return true
		}
	}
	return false
}
