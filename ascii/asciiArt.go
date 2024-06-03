package ascii

import (
	"fmt"
	"os"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
// The characters are extracted from a predefined file array.
// The function takes in four arguments: fileArr (a slice of strings representing the file array),
// wordsArr (a slice of strings representing the words to be printed),
// lettersToColor (a string representing the letters to be colored),
// and color (a string representing the color to be applied).
func Ascii(fileArr []string, wordsArr []string, lettersToColor string, colorCode string, outputfile string) {
	var count int
	reset := "\033[0m"
	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					switch {
					case len(lettersToColor) == 0:
						if IsFlagPassed("output") {
							fmt.Fprint(file, colorCode+fileArr[int(start)+i]+reset)
						} else {
							fmt.Print(colorCode + fileArr[int(start)+i] + reset)
						}
					case strings.Contains(lettersToColor, string(v)):
						if IsFlagPassed("output") {
							fmt.Fprintf(file, colorCode+fileArr[int(start)+i]+reset)
						} else {
							fmt.Print(colorCode + fileArr[int(start)+i] + reset)
						}
					default:
						if IsFlagPassed("output") {
							fmt.Fprintf(file, fileArr[int(start)+i])
						} else {
							fmt.Print(fileArr[int(start)+i])
						}
					}
				}
				if IsFlagPassed("output") {
					fmt.Fprintln(file)
				} else {
					fmt.Println()
				}
			}
		} else {
			count++
			if count < len(wordsArr) {
				if IsFlagPassed("output") {
					fmt.Fprintln(file)
				}
				fmt.Println()
			}
		}
	}
}
