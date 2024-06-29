package ascii

import (
	"errors"
	"strconv"
	"strings"
)

// Convert hex code to rgb code eg. #FFFFFF -> rgb(255, 255, 255)
// parameter: color an hex code and returns r, g, b uint8
func HexToRgb(color string) (r, g, b uint8, err error) {
	if strings.Contains(color, "=") {
		return 0, 0, 0, errors.New(" Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	}
	// Remove the '#' symbol if present
	color = strings.TrimPrefix(color, "#")

	// Split the hex code into red, green, and blue components
	var rStr, gStr, bStr string
	if len(color) == 6 {
		rStr = color[:2]
		gStr = color[2:4]
		bStr = color[4:]
	} else if len(color) == 3 {
		rStr = string(rune(color[0]))
		gStr = string(rune(color[1]))
		bStr = string(rune(color[2]))
		rStr += rStr
		gStr += gStr
		bStr += bStr
	}

	// parse each substring into an unsigned integer using base 16 (hexadecimal) and a bit size of 8 bits
	rInt, err := strconv.ParseUint(rStr, 16, 8)
	if err != nil {
		return 0, 0, 0, errors.New(" Error: Invalid Hex code")
	}
	gInt, err := strconv.ParseUint(gStr, 16, 8)
	if err != nil {
		return 0, 0, 0, errors.New(" Error: Invalid Hex code")
	}
	bInt, err := strconv.ParseUint(bStr, 16, 8)
	if err != nil {
		return 0, 0, 0, errors.New(" Error: Invalid Hex code")
	}

	// Return the red, green, and blue components as uint8 values
	return uint8(rInt), uint8(gInt), uint8(bInt), nil
}
