package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirectories(dirs ...string) error {
	for _, dir := range dirs {
		if err := CreateDirectory(dir); err != nil {
			return err
		}
	}
	return nil
}

func CreateDirectory(dirName string) error {
	return os.MkdirAll(dirName, os.ModePerm)
}

func CreateFile(name string) error {
	_, err := OpenFile(name)
	return err
}

func OpenFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_RDWR, os.ModePerm)
}

func AllFilesInDir(root string) []string {
	var files []string

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}
