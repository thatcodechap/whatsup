package utils

import (
	"regexp"
)

func RemoveMultipleSpaces(text string) string {
	return regexp.MustCompile(" {2,}").ReplaceAllString(text, " ")
}
