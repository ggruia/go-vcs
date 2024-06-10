package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

const (
	vcsRootDirName     = ".vcs"
	vcsCommitsDirPath  = ".vcs/objects"
	vcsBranchesDirPath = ".vcs/branches"
	vcsTimeFormat      = "2006-01-02 03:04:05"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "This allows you initialize vcs control system",
	Example: "vcs init",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runInitCommand()
	},
}

func runInitCommand() error {
	if exists, err := checkPathExists(vcsRootDirName); err != nil {
		return err
	} else if exists {
		return errors.New("vcs root directory is already exist")
	}

	if err := createDirectories(vcsRootDirName, vcsCommitsDirPath, vcsBranchesDirPath); err != nil {
		return err
	}

	fmt.Println("Repository initialized!")

	return nil
}
