package cli

import (
	"os"
	"path/filepath"
)

func addFile(filePath string) error {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Calculate the hash of the file contents
	hash := hashFileContents(contents)

	// Save the file contents in the objects directory
	objectPath := filepath.Join(vcsRootDir, commitsDirPath, hash)
	err = os.WriteFile(objectPath, contents, os.ModePerm)
	if err != nil {
		return err
	}

	// TODO: Update the index to include this file

	return nil
}
