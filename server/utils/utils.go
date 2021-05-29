package utils

import (
	"regexp"
)

func FindStringIndex(pattern string, s string) []int {
	re := regexp.MustCompile(pattern)
	return re.FindStringIndex(s)
}
