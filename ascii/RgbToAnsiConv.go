package ascii

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// RgbToAnsiConv converts an RGB color to an ANSI escape sequence.
// The function takes a string representing an RGB color in the format "rgb(R, G, B)"
// where R, G, and B are integers between 0 and 255.
// It returns a string that represents the ANSI escape sequence.
func RgbToAnsiConv(colorflag string) (string, error) {
	if strings.Contains(colorflag, "=") {
		return "", errors.New(" Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	}

	colorflag = strings.Trim(colorflag, "rgb")
	colorflag = strings.Trim(colorflag, "(){}[]")
	colorflag = strings.ReplaceAll(colorflag, " ", "")

	arr := strings.Split(colorflag, ",")
	if len(arr) != 3 {
		return "", errors.New("use rgb(r, g, b) format")
	}

	rgb := RGB{}
	for i, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return "", errors.New("the rgb code is wrong. use rgb(r, g, b) format")
		}

		if num >= 0 && num <= 255 {
			switch i {
			case 0:
				rgb.R = num
			case 1:
				rgb.G = num
			case 2:
				rgb.B = num
			default:
				return "", errors.New("the rgb code is wrong. use rgb(r, g, b) format")
			}
		} else {
			return "", errors.New("the RGB value should be between 0 and 255")
		}
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B), nil
}
