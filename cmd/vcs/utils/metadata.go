package utils

import (
	"bufio"
	"os"
	"strings"
)

const (
	separator  = " | "
	timeFormat = "2006-01-02 03:04:05"
)

type FileMetadata struct {
	Mtime string
	Path  string
	Work  string
	Stage string
	Repo  string
}

type FileMetadataMap map[string]FileMetadata

func parseMetadata(line string) FileMetadata {
	row := strings.Split(line, separator)
	return FileMetadata{
		Mtime: row[0],
		Path:  row[1],
		Work:  row[2],
		Stage: row[3],
		Repo:  row[4],
	}
}

func ReadMetadata(file string) (FileMetadataMap, error) {
	f, err := OpenFile(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	filesMetadata := FileMetadataMap{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		metadata := parseMetadata(scanner.Text())
		filesMetadata[metadata.Path] = metadata
	}

	return filesMetadata, nil
}

func (f *FileMetadata) formatMetadata() string {
	return strings.Join([]string{f.Mtime, f.Path, f.Work, f.Stage, f.Repo}, separator)
}

func WriteMetadata(file string, metadata FileMetadataMap) error {
	f, err := OpenFile(file)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, m := range metadata {
		line := m.formatMetadata()
		writer.WriteString(line + "\n")
	}

	writer.Flush()
	return nil
}

func UpdateMetadata(fileHashMap FileHashMap) FileMetadataMap {
	metadata := FileMetadataMap{}

	for p, f := range fileHashMap {
		file, err := os.ReadFile(f.Path)
		if err != nil {
			return nil
		}

		metadata[p] = FileMetadata{
			Mtime: f.Mtime.Format(timeFormat),
			Path:  f.Path,
			Work:  HashBytes(file),
			Stage: "-",
			Repo:  "-",
		}
	}

	return metadata
}
