package object

import (
	"fmt"
)

type Type string

const (
	BlobObject   Type   = "blob"
	TreeObject   Type   = "tree"
	CommitObject Type   = "commit"
	PathPrefix   string = ".vcs/objects/"
)

type Object struct {
	Type    Type
	Size    int
	Content []byte
}

func composeHeader(t Type, l int) []byte {
	return []byte(fmt.Sprintf("%s %d\n", t, l))
}

func createContent(o *Object) []byte {
	header := composeHeader(o.Type, o.Size)
	return append(header, o.Content...)
}
