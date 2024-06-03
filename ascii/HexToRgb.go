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

	color = strings.TrimPrefix(color, "#")

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

	return uint8(rInt), uint8(gInt), uint8(bInt), nil
}
