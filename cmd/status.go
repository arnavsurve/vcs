package cmd

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		// check .vcsignore
		ignorePatterns, err := ReadVCSIgnore()
		if err != nil {
			fmt.Println("Error reading .vcsignore file:", err)
		}

		// read the index file
		indexData, err := os.ReadFile(indexFile)
		if err != nil {
			fmt.Println("Error reading index file:", err)
			return
		}

		// unmarshal index data
		var index map[string]string
		if err := json.Unmarshal(indexData, &index); err != nil {
			fmt.Println("Error unmarshaling index data:", err)
			return
		}

		// display staged changes
		fmt.Println("Changes to be committed:")
		for file, hash := range index {
			fmt.Println(file, hash)
		}

		// display unstaged changes
		files, err := os.ReadDir(".")
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		// if file is not in index or hash is different, it has unstaged changes
		fmt.Println("\nUnstaged changes:")
		for _, file := range files {
			if file.IsDir() {
				continue // skip directories
			}
			filePath := file.Name()

			// check if file is ignored
			if IsIgnored(filePath, ignorePatterns) {
				continue
			}

			fileHash, err := computeFileHash(filePath)
			if err != nil {
				fmt.Println("Error computing hash for file:", filePath, err)
				continue
			}
			if indexHash, ok := index[filePath]; !ok || indexHash != fileHash {
				fmt.Println(filePath, "has changes")
			}
		}
	},
}

// compute SHA1 hash of a file
func computeFileHash(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	hash := sha1.Sum(content)
	return fmt.Sprintf("%x", hash), nil
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
