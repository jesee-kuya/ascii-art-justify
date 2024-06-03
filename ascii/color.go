package ascii

import (
	"fmt"
	"strings"
)

// Color takes words a slice of string from a certain index, joins the words,
// parameters: colorflag string, lettersTocolor string, argsPassed []string, bannerContent []string
// and calls the Ascii function to print the words in ascii-art
func Color(colorflag string, lettersTocolor string, argsPassed []string, bannerContent []string, outputfile string) {
	var str string
	var err error
	var colorCode string
	var rgb RGB
	colorflag = strings.ToLower(colorflag)

	_, ok := Colormap[colorflag]
	if !ok {
		if strings.Contains(colorflag, "rgb") {
			str, err = RgbToAnsiConv(colorflag)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if strings.Contains(colorflag, "#") {
			r, g, b, err := HexToRgb(colorflag)
			if err != nil {
				fmt.Println(err)
				return
			}

			rgb.R = int(r)
			rgb.G = int(g)
			rgb.B = int(b)
			str = fmt.Sprintf("\033[38;2;%v;%v;%vm", rgb.R, rgb.G, rgb.B)
		} else {
			if strings.Contains(colorflag, "=") {
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			} else {
				fmt.Printf("The color %v is not yet defined. Try another color.\n", colorflag)
			}
			return
		}
	}

	if str == "" {
		colorCode = Colormap[colorflag]
	} else {
		colorCode = str
	}

	if len(argsPassed) == 1 {
		Art(argsPassed, bannerContent, lettersTocolor, colorCode, 0, outputfile)
	} else if len(argsPassed) == 2 {
		lettersTocolor = argsPassed[0]
		Art(argsPassed, bannerContent, lettersTocolor, colorCode, 1, outputfile)
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
}
