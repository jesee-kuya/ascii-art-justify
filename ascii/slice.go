package ascii

import (
	"regexp"
	"strings"
)

// Slice splits a string at \n and returns a slice of strings
func Slice(word string) []string {
	// Add a tab before the next word
	re := regexp.MustCompile(`\\t`)
	word = re.ReplaceAllString(word, "    $0")
	word = re.ReplaceAllString(word, "")

	// Add a tab and a new line before the next word
	re1 := regexp.MustCompile(`\\v|\\f`)
	word = re1.ReplaceAllString(word, "$0\\n    ")
	word = re1.ReplaceAllString(word, "")

	// Replace the previous words or characters
	re2 := regexp.MustCompile(`(.*?)(\\r)`)
	word = re2.ReplaceAllString(word, "")

	// Delete a character before or after
	re3 := regexp.MustCompile(`.\\b|DEL.`)
	word = re3.ReplaceAllString(word, "")

	// Handle new line, both \n and \\n
	re4 := regexp.MustCompile(`\\n`)
	word = re4.ReplaceAllLiteralString(word, "\n")
	wordArr := strings.Split(word, "\n")

	return wordArr
}
