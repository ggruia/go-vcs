package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Commit struct {
	ID        string
	Timestamp time.Time
	Message   string
	Files     map[string]string
	ParentID  string
}

func CreateCommit(message string) error {
	index, err := readIndex()
	if err != nil {
		return err
	}

	commit := Commit{
		ID:        "",
		Timestamp: time.Now(),
		Message:   message,
		Files:     index,
		ParentID:  "",
	}

	// serialize the commit object
	commitData, err := json.Marshal(commit)
	if err != nil {
		return err
	}

	// generate a hash for the commit
	commit.ID = hashFileContents(commitData)

	// write the commit object to the commits directory
	commitPath := filepath.Join(VcsDir, CommitsDir, commit.ID)
	err = os.WriteFile(commitPath, commitData, os.ModePerm)
	if err != nil {
		return err
	}

	// update the HEAD to point to the new commit
	headPath := filepath.Join(VcsDir, "HEAD")
	return os.WriteFile(headPath, []byte(commit.ID), os.ModePerm)
}
