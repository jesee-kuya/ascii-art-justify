package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
// The characters are extracted from a predefined file array.
// The function takes in four arguments: fileArr (a slice of strings representing the file array),
// wordsArr (a slice of strings representing the words to be printed),
// lettersToColor (a string representing the letters to be colored),
// and color (a string representing the color to be applied).
func AsciiAlign(fileArr []string, wordsArr []string, lettersToColor string, colorCode string, alignFlag string) {
	var toPrint [10]string
	var str string
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	size, err := cmd.Output()
	fmt.Println(string((size)))
	if err != nil {
		fmt.Println(err)
		return
	}
	widtharr := strings.Split(strings.ReplaceAll(string(size), "\n", ""), " ")
	termWidth, err := strconv.Atoi(widtharr[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var count int
	reset := "\033[0m"

	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					switch {
					case len(lettersToColor) == 0:
						// fmt.Printf(fmt.Sprintf("%%-%ds", termWidth/2), fmt.Sprintf(fmt.Sprintf("%%%ds", termWidth/2),
						toPrint[i] += colorCode + fileArr[int(start)+i] + reset
					case strings.Contains(lettersToColor, string(v)):
						toPrint[i] += colorCode + fileArr[int(start)+i] + reset
					default:
						toPrint[i] += fileArr[int(start)+i]
					}
				}
			}
		} else {
			count++
			if count < len(wordsArr) {
				fmt.Printf(fmt.Sprintf("%%-%ds", (termWidth/2)), fmt.Sprintf(fmt.Sprintf("%%%ds", termWidth/2), "\n"))
			}
		}
		for _, v := range toPrint {
			if v != "" {
				if alignFlag == "centre" {
					str = fmt.Sprintf("%*s", -termWidth, fmt.Sprintf("%*s", (termWidth+len(v))/2, v))
				} else if alignFlag == "right" {
					str = fmt.Sprintf("%*s", +(termWidth), v)
				} else if alignFlag == "left" {
					str = fmt.Sprintf("%v", v)
				}
				fmt.Println(str)
			}
		}
		toPrint = [10]string{}
	}
}
