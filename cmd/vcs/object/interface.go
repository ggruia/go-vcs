package object

import (
	"bytes"
	"fmt"
	"go-vcs/cmd/vcs/utils"
	"strconv"
	"strings"
)

type ObjReader interface {
	Read(path string) (*Object, error)
}

type ObjWriter interface {
	Write(object *Object) error
}

type ObjManager struct {
	Path string
}

func (manager *ObjManager) Read(name string) (*Object, error) {
	data, err := utils.ReadFile(manager.Path + name)
	if err != nil {
		return nil, err
	}

	b, err := utils.DecompressBytes(data)
	if err != nil {
		return nil, err
	}

	h, c, ok := bytes.Cut(b, []byte("\n"))
	if !ok {
		return nil, fmt.Errorf("invalid format")
	}

	header := string(h)
	headerComponents := strings.Split(header, " ")
	size, err := strconv.Atoi(headerComponents[1])
	if err != nil {
		return nil, err
	}

	o := &Object{
		Type:    Type(headerComponents[0]),
		Size:    size,
		Content: c,
	}

	return o, nil
}

func (manager *ObjManager) Write(o *Object) error {
	content := createContent(o)
	path := manager.Path + utils.HashBytes(content)

	return utils.WriteCompressed(path, content)
}
