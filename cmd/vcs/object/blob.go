package object

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go-vcs/cmd/vcs/utils"
)

type Blob struct {
	ID   string
	Size int
	Data []byte
}

func (b *Blob) GetID() string {
	return b.ID
}

func (b *Blob) GetType() Type {
	return blobObject
}

func (b *Blob) GetData() []byte {
	return b.Data
}

func (b *Blob) GetSize() int {
	return b.Size
}

func (b *Blob) ToBytes() ([]byte, error) {
	header := fmt.Sprintf("%s %d\n", blobObject, b.Size)

	var buf bytes.Buffer
	buf.Write([]byte(header))
	buf.Write(b.Data)

	return buf.Bytes(), nil
}

func BytesToBlob(b []byte) (*Blob, error) {
	header, data, ok := bytes.Cut(b, []byte("\n"))
	if !ok {
		return nil, errors.New("invalid format")
	}

	t, s, ok := bytes.Cut(header, []byte(" "))
	if !ok {
		return nil, errors.New("invalid format")
	}

	if Type(t) != blobObject {
		return nil, errors.New("invalid type")
	}

	size := int(binary.BigEndian.Uint64(s))
	if size != len(data) {
		return nil, errors.New("invalid size")
	}

	blob := &Blob{
		ID:   utils.HashBytes(b),
		Size: size,
		Data: data,
	}

	return blob, nil
}

func NewBlob(data []byte) *Blob {
	content := CreateContent(data, blobObject)
	return &Blob{
		ID:   utils.HashBytes(content),
		Size: len(data),
		Data: data,
	}
}
