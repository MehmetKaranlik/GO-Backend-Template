package Helpers

import "strings"

func GetFileExtension(fileName string) string {
	parts := strings.Split(fileName, ".")
	splintedLength := len(parts)
	last := parts[splintedLength-1]
	return last
}

func GetFileName(fileName string) string {
	parts := strings.Split(fileName, ".")
	return strings.Join(parts[:len(parts)-1], ".")
}
