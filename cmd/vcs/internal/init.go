package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo() error {
	if _, err := os.Stat(VcsDir); !os.IsNotExist(err) {
		return fmt.Errorf("repository already exists")
	}

	// create the .vcs directory and subdirectories
	err := os.MkdirAll(filepath.Join(VcsDir, ObjectsDir), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(VcsDir, ObjectsDir), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
