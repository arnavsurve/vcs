package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit -m [message]",
	Short: "Commit changes to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")
		if message == "" {
			fmt.Println("Commit message is required")
			return
		}
		commit(message)
		fmt.Println("Changes committed with message:", message)
	},
}

func commit(message string) {
	indexData, _ := os.ReadFile(indexFile)

	commit := map[string]interface{}{
		"message":   message,
		"timestamp": time.Now().Format(time.RFC3339),
		"index":     string(indexData),
	}

	// hash commit data and convert to string
	commitData, _ := json.Marshal(commit)
	commitHash := sha1.New()
	commitHash.Write(commitData)
	commitSha1 := hex.EncodeToString(commitHash.Sum(nil))

	commitPath := filepath.Join(objectsDir, commitSha1)
	os.WriteFile(commitPath, commitData, 0644)

	// update HEAD with the new commit's SHA1
	headFile := filepath.Join(repoDir, "HEAD")
	os.WriteFile(headFile, []byte(commitSha1), 0644)
}

func init() {
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	rootCmd.AddCommand(commitCmd)
}
