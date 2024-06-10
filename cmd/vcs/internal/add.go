package internal

import (
	"os"
	"path/filepath"
)

func AddFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// hash of the file contents
	hash := hashFileContents(data)

	// save file contents in the objects directory
	objectPath := filepath.Join(VcsDir, ObjectsDir, hash)
	err = os.WriteFile(objectPath, data, os.ModePerm)
	if err != nil {
		return err
	}

	err = updateIndex(objectPath, hash)
	if err != nil {
		return err
	}

	return nil
}
