package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const indexFile = "index"

func updateIndex(filePath string, hash string) error {
	indexFilePath := filepath.Join(VcsDir, indexFile)
	indexContents := fmt.Sprintf("%s %s\n", filePath, hash)

	// Append the new file and its hash to the index
	return os.WriteFile(indexFilePath, []byte(indexContents), os.ModeAppend)
}

func readIndex() (map[string]string, error) {
	index := make(map[string]string)
	indexFilePath := filepath.Join(VcsDir, indexFile)
	contents, err := os.ReadFile(indexFilePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid index line: %s", line)
		}
		index[parts[0]] = parts[1]
	}

	return index, nil
}
