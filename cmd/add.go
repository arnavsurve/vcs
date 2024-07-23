package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a file to the repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ignorePatterns, err := ReadVCSIgnore()
		if err != nil {
			fmt.Println("Error reading .vcsignore file:", err)
			return
		}

		// check if file is ignored
		filePath := args[0]
		if IsIgnored(filePath, ignorePatterns) {
			fmt.Println("File is ignored:", filePath)
			return
		}

		addFile(filePath)
		fmt.Println("File added:", filePath)
	},
}

// add file to the index
func addFile(filePath string) {
	// convert indexFile to SHA1
	content, _ := os.ReadFile(indexFile)
	hash := sha1.New()
	hash.Write(content)
	sha1Sum := hex.EncodeToString(hash.Sum(nil)) // encode SHA1 hash to string

	objectPath := filepath.Join(objectsDir, sha1Sum)
	os.WriteFile(objectPath, content, 0644)

	// parse index data and update index with new file
	indexData, _ := os.ReadFile(indexFile)
	var index map[string]string
	json.Unmarshal(indexData, &index)
	index[filePath] = sha1Sum

	// convert index back to JSOn and write to index file
	newIndexData, _ := json.Marshal(index)
	os.WriteFile(indexFile, newIndexData, 0644)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
