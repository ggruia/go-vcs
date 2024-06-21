package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type FileStat struct {
	Name  string
	Size  int64
	Data  []byte
	Mtime time.Time
}

type StatMap map[string]FileStat

var excludedPrefixes = []string{".git", ".vcs", "vcs", ".idea"}

func ReadFilesFromWorkingDir(rootDir string) (StatMap, error) {
	fileMap := StatMap{}

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

		if StartsWithAny(relPath, excludedPrefixes) {
			return nil
		}

		data, err := os.ReadFile(relPath)
		if err != nil {
			return err
		}

		stat, err := os.Stat(relPath)
		if err != nil {
			return err
		}

		fileMap[relPath] = FileStat{
			Name:  stat.Name(),
			Size:  stat.Size(),
			Data:  data,
			Mtime: stat.ModTime(),
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileMap, nil
}
