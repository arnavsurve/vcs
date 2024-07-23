package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

const repoDir = ".vcs"
const objectsDir = ".vcs/objects"
const indexFile = ".vcs/index"

// init initializes the repository
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(repoDir); os.IsNotExist(err) {
			os.Mkdir(repoDir, 0755)
			os.Mkdir(objectsDir, 0755)
			initIndex()
			fmt.Println("Initialized empty vcs repository in", repoDir)
		} else {
			fmt.Println("Repository already exists")
		}
	},
}

func initIndex() {
	index := make(map[string]string)
	indexData, _ := json.Marshal(index)
	os.WriteFile(indexFile, indexData, 0644)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
