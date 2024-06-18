package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"go-vcs/cmd/vcs/utils"
	"path/filepath"
)

const (
	vcsRootDir     = ".vcs"
	vcsObjectsDir  = "objects"
	vcsBranchesDir = "branches"
	vcsRefsDir     = "refs"
	vcsIndexPath   = "index"
	vcsHeadPath    = "HEAD"
	vcsTimeFormat  = "2006-01-02 03:04:05"
)

var (
	objectsDirPath  = filepath.Join(vcsRootDir, vcsObjectsDir)
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
	if exists, err := utils.CheckPathExists(vcsRootDir); err != nil {
		return err
	} else if exists {
		return errors.New("vcs root directory already exists")
	}

	if err := utils.CreateDirectories(vcsRootDir, objectsDirPath, branchesDirPath, refsDirPath); err != nil {
		return err
	}

	if err := utils.CreateFile(indexPath); err != nil {
		return err
	}

	if err := utils.CreateFile(headPath); err != nil {
		return err
	}

	fmt.Println("Repository initialized!")

	return nil
}
