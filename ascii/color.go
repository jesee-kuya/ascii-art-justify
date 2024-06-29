package ascii

import (
	"fmt"
	"strings"
)

// Color takes words a slice of string from a certain index, joins the words,
// parameters: colorflag string, lettersTocolor string, argsPassed []string, bannerContent []string
// and calls the Ascii function to print the words in ascii-art
func (s *Receiver) Color() {
	var str string
	var err error
	var rgb RGB
	colorflag := strings.ToLower(s.Colorflag)

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
		s.ColorCode = Colormap[colorflag]
	} else {
		s.ColorCode = str
	}

	if len(s.ArgsPassed) == 1 {
		s.IndexToStartDisplay = 0
	} else if len(s.ArgsPassed) == 2 {
		s.LettersToColor = s.ArgsPassed[0]
		s.IndexToStartDisplay = 1
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
}
