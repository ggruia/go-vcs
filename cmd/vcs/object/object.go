package object

import (
	"fmt"
)

type Type string

type Object interface {
	GetID() string
	GetType() Type
	GetData() []byte
	GetSize() int
	ToBytes() ([]byte, error)
}

const (
	blobObject   Type   = "blob"
	treeObject   Type   = "tree"
	commitObject Type   = "commit"
	pathPrefix   string = ".vcs/objects/"
)

func composeHeader(t Type, l int) []byte {
	return []byte(fmt.Sprintf("%s %d\n", t, l))
}

func CreateContent(data []byte, t Type) []byte {
	header := composeHeader(t, len(data))
	return append(header, data...)
}
