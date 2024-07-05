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
func (s *Receiver) GetSizeOfCharacters(word string) int {
	var length int
	for _, v := range word {
		start := ((v - 32) * 9) + 4
		length += len(s.FileArr[start])
	}
	return length
}

// SizeOfSpace returns the width of ascii-art space as an int
func (s *Receiver) SizeOfSpace() int {
	v := ' '
	start := ((v - 32) * 9) + 4
	return len(s.FileArr[start])
}

func CheckSpace(word string) (check int) {
	for _, v := range word {
		if v == ' ' {
			check++
		}
	}
	return
}

// AddSpace adds spaces in between words and returns the new string with the added spaces
func (s *Receiver) AddSpace(word string, space int) (new string) {
	var sp string
	str := word
	var count int
	width := Getwidth()
	for s.GetSizeOfCharacters(str)+s.SizeOfSpace() < width {
		str += " "
		count++
	}
	target := count / space

	// for target%s.SizeOfSpace() != 0 {
	// 	target--
	// }

	for len(sp) != target {
		sp += " "
	}

	for _, v := range word {
		if s.GetSizeOfCharacters(new) < width {
			if v == ' ' {
				new += sp
			}
			new += string(v)
		}
	}
	return
}

func (s *Receiver) Allign(n int) {
	var arr []string
	width := Getwidth()
	for _, v := range s.WordsArr {
		if v == "" {
			arr = append(arr, v)
			continue
		} else {
			lenWord := s.GetSizeOfCharacters(v)
			target := width/n - lenWord/n
			space := ""

			for target%s.SizeOfSpace() != 0 {
				target--
			}

			for s.GetSizeOfCharacters(space) != target {
				space += " "
			}
			v = space + v
			arr = append(arr, v)
		}
	}
	s.WordsArr = arr

}

// AllignRight adds the necessary spaces before the string to align its ascii right on the terminal
func (s *Receiver) AllignRight() {
	s.Allign(1)
}

// AllignCentre adds the necessary spaces so that the text can be printed centre of the terminal
func (s *Receiver) AllignCentre() {
	s.Allign(2)
}

// AllignJustify adds necessary spaces in between words to justify align the ascii-art on the terminal
func (s *Receiver) AllignJustify() {
	var arr []string
	for _, v := range s.WordsArr {
		if v == "" {
			arr = append(arr, v)
			continue
		} else {
			spaces := CheckSpace(v)
			if spaces != 0 {
				v = s.AddSpace(v, spaces)
			}
			arr = append(arr, v)
		}
	}
	s.WordsArr = arr
}
