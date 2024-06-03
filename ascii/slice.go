package ascii

import (
	"regexp"
	"strings"
)

// Slice splits a string at \n and returns a slice of strings
func Slice(word string) []string {
	re := regexp.MustCompile(`\\t`)
	word = re.ReplaceAllString(word, "    $0")
	word = re.ReplaceAllString(word, "")

	re1 := regexp.MustCompile(`\\v|\\f`)
	word = re1.ReplaceAllString(word, "$0\\n    ")
	word = re1.ReplaceAllString(word, "")

	re2 := regexp.MustCompile(`(.*?)(\\r)`)
	word = re2.ReplaceAllString(word, "")

	re3 := regexp.MustCompile(`.\\b|DEL.`)
	word = re3.ReplaceAllString(word, "")

	re4 := regexp.MustCompile(`\\n`)
	word = re4.ReplaceAllLiteralString(word, "\n")
	wordArr := strings.Split(word, "\n")

	return wordArr
}
