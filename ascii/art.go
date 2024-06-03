package ascii

// Art takes words a slice of string from a certain index, joins the words,
// checks if there is a non-ascii character
// and calls the Ascii function to print the words in ascii-art
func Art(argsPassed []string, content []string, lettersTocolor string, colorCode string, indexToStartDisplay int, outputfile string) {
	word := Arrange(argsPassed[indexToStartDisplay:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(content, wordsArr, lettersTocolor, colorCode, outputfile)
}
