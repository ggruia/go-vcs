package object

import (
	"bufio"
	"go-vcs/cmd/vcs/utils"
	"reflect"
	"sort"
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
	AddToStaging(files []string)
}

type MetadataManager struct {
	Path string
}

func (manager *MetadataManager) Read() (MetadataMap, error) {
	f, err := utils.OpenFile(manager.Path)
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
	f, err := utils.OpenFile(manager.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, k := range sortedMapKeys(*metadataMap) {
		line := (*metadataMap)[k].formatMetadata()
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

	statMap, err := utils.ReadFilesFromWorkingDir(".")
	if err != nil {
		return err
	}

	for path, stat := range statMap {
		mTime := stat.Mtime.Format(timeFormat)
		blob := NewBlob(stat.Data)
		if metadata, ok := metadataMap[path]; ok {
			if metadata.Work == blob.ID {
				continue
			}

			metadata.Work = blob.ID
			metadata.Mtime = mTime
			metadataMap[path] = metadata
		} else {
			metadataMap[path] = Metadata{
				Mtime: mTime,
				Path:  path,
				Work:  blob.ID,
				Stage: empty,
				Repo:  empty,
			}
		}
	}

	if err := manager.Write(&metadataMap); err != nil {
		return err
	}

	return nil
}

func (manager *MetadataManager) AddToStaging(files []string) {
	metadataMap, err := manager.Read()
	if err != nil {
		return
	}

	for _, file := range files {
		if metadata, ok := metadataMap[file]; ok {
			metadata.Stage = metadata.Work
			metadataMap[file] = metadata
		}
	}

	if err := manager.Write(&metadataMap); err != nil {
		return
	}
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

func (m Metadata) formatMetadata() string {
	return strings.Join([]string{m.Mtime, m.Path, m.Work, m.Stage, m.Repo}, separator)
}

func sortedMapKeys(m interface{}) (keys []string) {
	mapKeys := reflect.ValueOf(m).MapKeys()

	for _, key := range mapKeys {
		keys = append(keys, key.Interface().(string))
	}
	sort.Strings(keys)
	return
}
