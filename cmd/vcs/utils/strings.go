package utils

import "strings"

func StartsWithAny(s string, prefixes []string) bool {
	for _, l := range prefixes {
		if strings.HasPrefix(s, l) {
			return true
		}
	}
	return false
}
