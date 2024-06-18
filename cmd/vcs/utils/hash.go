package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashBytes(contents []byte) string {
	hasher := sha1.New()
	hasher.Write(contents)
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkIdenticalHashes(currentHash, targetHash string) bool {
	return currentHash == targetHash
}
