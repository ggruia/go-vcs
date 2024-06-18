package utils

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func compressBytes(data []byte) ([]byte, error) {
	var compressedBuffer bytes.Buffer
	zlibWriter := zlib.NewWriter(&compressedBuffer)

	_, err := zlibWriter.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to write data to zlib writer: %w", err)
	}

	err = zlibWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close zlib writer: %w", err)
	}

	return compressedBuffer.Bytes(), nil
}

func DecompressBytes(data []byte) ([]byte, error) {
	buffer := bytes.NewReader(data)
	zlibReader, err := zlib.NewReader(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to create zlib reader: %w", err)
	}
	defer zlibReader.Close()

	var decompressedBuffer bytes.Buffer
	_, err = io.Copy(&decompressedBuffer, zlibReader)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}

	return decompressedBuffer.Bytes(), nil
}
