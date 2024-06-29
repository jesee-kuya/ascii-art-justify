package ascii

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Getwidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	sizeArr := strings.Fields(string(out))
	width, _ := strconv.Atoi(sizeArr[1])
	return width
}

func (s *Receiver) GetSizeOfCharacters(word string) int {
	var length int
	for _, v := range word {
		start := ((v - 32) * 9) + 4
		length += len(s.FileArr[start])
	}
	return length
}

func (s *Receiver) SizeOfSpace() int {
	v := ' '
	start := ((v - 32) * 9) + 4
	return len(s.FileArr[start])
}

func (s *Receiver) AllignCentre() {
	var checkLen int
	var word string
	for _, v := range s.WordsArr {
		if len(v) > checkLen {
			word = v
			checkLen = len(v)
		}
	}
	lenChar := s.GetSizeOfCharacters(word)
	rem := lenChar % 2

	width := Getwidth()
	var arr []string
	space := " "
	for len(space)*s.SizeOfSpace() <= (width/2 - ((lenChar / 2) - rem)) {
		space += " "
	}
	for _, v := range s.WordsArr {
		word := v
		word = space + word
		arr = append(arr, word)
	}
	s.WordsArr = arr
}
