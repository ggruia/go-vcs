package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Checkout(commitID string) error {
	commitPath := filepath.Join(VcsDir, ObjectsDir, commitID)
	commitData, err := os.ReadFile(commitPath)
	if err != nil {
		return err
	}

	var commit Commit
	err = json.Unmarshal(commitData, &commit)
	if err != nil {
		return err
	}

	for filePath, hash := range commit.Files {
		objectPath := filepath.Join(VcsDir, ObjectsDir, hash)
		fileContents, err := os.ReadFile(objectPath)
		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, fileContents, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
