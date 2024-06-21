package object

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go-vcs/cmd/vcs/utils"
)

type Reference struct {
	ID   string
	Type Type
	Name string
}

type Tree struct {
	ID         string
	Size       int
	References []Reference
}

func (t *Tree) GetID() string {
	return t.ID
}

func (t *Tree) GetType() Type {
	return treeObject
}

func (t *Tree) GetData() []byte {
	var buf bytes.Buffer

	for _, ref := range t.References {
		line := fmt.Sprintf("%s %s %s\n", ref.Type, ref.ID, ref.Name)
		buf.Write([]byte(line))
	}

	return buf.Bytes()
}

func (t *Tree) GetSize() int {
	return t.Size
}

func (t *Tree) ToBytes() ([]byte, error) {
	header := fmt.Sprintf("%s %d\n", treeObject, t.Size)

	var buf bytes.Buffer
	buf.Write([]byte(header))
	buf.Write(t.GetData())

	return buf.Bytes(), nil
}

func BytesToTree(b []byte) (*Tree, error) {
	header, data, ok := bytes.Cut(b, []byte("\n"))
	if !ok {
		return nil, errors.New("invalid format")
	}

	t, s, ok := bytes.Cut(header, []byte(" "))
	if !ok {
		return nil, errors.New("invalid format")
	}

	if Type(t) != treeObject {
		return nil, errors.New("invalid type")
	}

	size := int(binary.BigEndian.Uint64(s))
	if size != len(data) {
		return nil, errors.New("invalid size")
	}

	var references []Reference
	for _, ref := range bytes.Split(data, []byte("\n")) {
		refComponents := bytes.Split(ref, []byte(" "))
		reference := Reference{
			Type: Type(refComponents[0]),
			ID:   string(refComponents[1]),
			Name: string(refComponents[2]),
		}
		references = append(references, reference)
	}

	tree := &Tree{
		ID:         utils.HashBytes(b),
		Size:       size,
		References: references,
	}

	return tree, nil
}
