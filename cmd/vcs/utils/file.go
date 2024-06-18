package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file: %w", err)
	}

	return data, nil
}

func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func WriteCompressed(path string, data []byte) error {
	compressedData, err := compressBytes(data)
	if err != nil {
		return fmt.Errorf("failed to compress bytes: %w", err)
	}

	err = WriteFile(path, compressedData)
	if err != nil {
		return err
	}

	return nil
}
