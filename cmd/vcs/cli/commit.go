package cli

import (
	"errors"
	"github.com/spf13/cobra"
)

const (
	messageFlag = "message"
)

func init() {
	commitCmd.Flags().StringP(messageFlag, "m", "", "the commit message")
	_ = commitCmd.MarkFlagRequired(messageFlag)
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "This allows you to commit tracked file changes",
	Example: "vcs commit -m \"commit message\"",
	RunE: func(cmd *cobra.Command, args []string) error {
		msg, _ := cmd.Flags().GetString(messageFlag)
		if msg == "" {
			return errors.New("commit message cannot be empty")
		}
		return runCommitCommand(indexPath, msg)
	},
}

func runCommitCommand(stagingPath string, message string) error {
	//todo

	return nil
}
