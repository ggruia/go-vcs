package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileStat struct {
	Path  string
	Hash  string
	Size  int64
	Mtime time.Time
}

type FileHashMap map[string]FileStat

var ExcludedPrefixes = []string{".git", ".vcs", "vcs"}

func ReadFilesFromWorkingDir(rootDir string) (FileHashMap, error) {
	fileMap := FileHashMap{}

	err := filepath.WalkDir(rootDir, func(path string, fsEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if fsEntry.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		for _, prefix := range ExcludedPrefixes {
			if strings.HasPrefix(relPath, prefix) {
				return nil
			}
		}

		data, err := os.ReadFile(relPath)
		if err != nil {
			return err
		}

		fileInfo, err := os.Stat(relPath)
		if err != nil {
			return err
		}

		fileMap[relPath] = FileStat{
			Path:  relPath,
			Hash:  HashBytes(data),
			Size:  fileInfo.Size(),
			Mtime: fileInfo.ModTime(),
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileMap, nil
}
