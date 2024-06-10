package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	statusFilePath      = filepath.Join(vcsRootDirName, "status.txt")
	stagingAreaFilePath = filepath.Join(vcsRootDirName, "staging-area.txt")
)

const (
	vcsRootDirName            = ".vcs"
	vcsCommitDirPath          = ".vcs/commit"
	vcsCheckoutDirPath        = ".vcs/checkout"
	vcsStatusFilePath         = ".vcs/status.txt"
	vcsStagingFilePath        = ".vcs/staging-area.txt"
	vcsTimeFormat             = "2006-01-02 03:04:05"
	vcsCommitMetadataFileName = "metadata.txt"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "This allows you initialize vX control system",
	Example: "vx init",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runInitCommand()
	},
}

func runInitCommand() error {
	if exists, err := checkPathExists(vcsRootDirName); err != nil {
		return err
	} else if exists {
		return errors.New("vx root directory is already exist")
	}

	if err := createDirectories(vcsRootDirName, vcsCommitDirPath, vcsCheckoutDirPath); err != nil {
		return err
	}

	if err := createFile(vcsStatusFilePath); err != nil {
		return err
	}

	if err := createFile(vcsStagingFilePath); err != nil {
		return err
	}

	fmt.Println("Repository initialized!")

	return nil
}
