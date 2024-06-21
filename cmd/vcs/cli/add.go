package cli

import (
	"github.com/spf13/cobra"
	"go-vcs/cmd/vcs/object"
	"go-vcs/cmd/vcs/utils"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "This allows you to track all file status (created, modified)",
	Example: "vcs add main.go images/",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAddCommand(indexPath, args)
	},
}

func runAddCommand(stagingPath string, filePaths []string) error {
	manager := object.MetadataManager{Path: stagingPath}

	for _, path := range filePaths {
		files := utils.AllFilesInDir(path)
		manager.AddToStaging(files)
	}

	return nil
}
