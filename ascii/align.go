package ascii

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Getwidth gets the width of the current terminal and return it as an int
func Getwidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	sizeArr := strings.Fields(string(out))
	width, _ := strconv.Atoi(sizeArr[1])
	return width
}

// GetSizeOfCharacters return the total length of the charactes on the terminal in ascii-art format as an int
func (s *Receiver) RemainingSpace(word string) {
	var length int
	for _, v := range word {
		start := ((v - 32) * 9) + 4
		length += len(s.FileArr[start])
	}
	s.Spaces = append(s.Spaces, Getwidth()-length)
}

func (s *Receiver) Alligner() {
	for _, v := range s.WordsArr {
		s.RemainingSpace(v)
	}
}