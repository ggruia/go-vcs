package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

const (
	vcsRootDir     = ".vcs"
	vcsCommitsDir  = "objects"
	vcsBranchesDir = "branches"
	vcsRefsDir     = "refs"
	vcsIndexPath   = "index"
	vcsHeadPath    = "HEAD"
	vcsTimeFormat  = "2006-01-02 03:04:05"
)

var (
	commitsDirPath  = filepath.Join(vcsRootDir, vcsCommitsDir)
	branchesDirPath = filepath.Join(vcsRootDir, vcsBranchesDir)
	refsDirPath     = filepath.Join(vcsRootDir, vcsRefsDir)
	indexPath       = filepath.Join(vcsRootDir, vcsIndexPath)
	headPath        = filepath.Join(vcsRootDir, vcsHeadPath)
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
	if exists, err := checkPathExists(vcsRootDir); err != nil {
		return err
	} else if exists {
		return errors.New("vcs root directory is already exist")
	}

	if err := createDirectories(vcsRootDir, commitsDirPath, branchesDirPath, refsDirPath); err != nil {
		return err
	}

	if err := createFile(indexPath); err != nil {
		return err
	}

	if err := createFile(headPath); err != nil {
		return err
	}

	fmt.Println("Repository initialized!")

	return nil
}
