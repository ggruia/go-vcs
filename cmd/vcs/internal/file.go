package internal

import (
	"crypto/sha1"
	"encoding/hex"
)

const (
	VcsDir     = ".vcs"    // VCS metadata
	ObjectsDir = "objects" // object files
	CommitsDir = "commits" // commit files
)

func hashFileContents(contents []byte) string {
	hasher := sha1.New()
	hasher.Write(contents)
	return hex.EncodeToString(hasher.Sum(nil))
}
