package utils

import (
	"bufio"
	"go-vcs/cmd/vcs/object"
	"strings"
)

const (
	separator  = " | "
	timeFormat = "2006-01-02 03:04:05"
	empty      = "-"
)

type Metadata struct {
	Mtime string
	Path  string
	Work  string
	Stage string
	Repo  string
}

type MetadataMap map[string]Metadata

type MetadataIO interface {
	Read() (MetadataMap, error)
	Write(metadataMap *MetadataMap) error
	UpdateFromWorkDir() error
}

type MetadataManager struct {
	Path string
}

func (manager *MetadataManager) Read() (MetadataMap, error) {
	f, err := OpenFile(manager.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	metadataMap := MetadataMap{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		metadata := parseMetadata(scanner.Text())
		metadataMap[metadata.Path] = metadata
	}

	return metadataMap, nil
}

func (manager *MetadataManager) Write(metadataMap *MetadataMap) error {
	f, err := OpenFile(manager.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, m := range *metadataMap {
		line := m.formatMetadata()
		writer.WriteString(line + "\n")
	}

	writer.Flush()
	return nil
}

func (manager *MetadataManager) UpdateFromWorkDir() error {
	metadataMap, err := manager.Read()
	if err != nil {
		return err
	}

	statMap, err := ReadFilesFromWorkingDir(".")
	if err != nil {
		return err
	}

	for path, stat := range statMap {
		mTime := stat.Mtime.Format(timeFormat)
		if metadata, ok := metadataMap[path]; ok {
			blob := object.NewBlob(stat.Data)
			metadata.Work = blob.ID
			metadata.Mtime = mTime
		} else {
			metadataMap[path] = Metadata{
				Mtime: mTime,
				Path:  path,
				Work:  HashBytes(stat.Data),
				Stage: empty,
				Repo:  empty,
			}
		}
	}

	return nil
}

func parseMetadata(line string) Metadata {
	row := strings.Split(line, separator)
	return Metadata{
		Mtime: row[0],
		Path:  row[1],
		Work:  row[2],
		Stage: row[3],
		Repo:  row[4],
	}
}

func (f *Metadata) formatMetadata() string {
	return strings.Join([]string{f.Mtime, f.Path, f.Work, f.Stage, f.Repo}, separator)
}
