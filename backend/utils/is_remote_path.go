package utils

import "strings"

func IsRemotePath(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
