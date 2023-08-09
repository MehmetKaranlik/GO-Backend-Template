package Methods

import (
	"strconv"
	"strings"
	"time"
)

func GetExtension(fileName string) string {
	transformed := strings.Split(fileName, ".")
	if len(transformed) <= 1 {
		return ""
	}
	return transformed[len(transformed)-1]
}

func GetFilenameWithoutExtension(fileName string) string {
	transformed := strings.Split(fileName, ".")
	if len(transformed) <= 1 {
		return ""
	}
	return strings.Join(transformed[:len(transformed)-1], ".")
}

func ConstructNewNameForFile(fileName string) string {
	baseName := GetFilenameWithoutExtension(fileName)
	extension := GetExtension(fileName)
	newName := baseName + strconv.FormatInt(time.Now().UnixNano(), 10)
	return newName + "." + extension
}
