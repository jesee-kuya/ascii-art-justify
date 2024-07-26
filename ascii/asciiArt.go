package ascii

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Ascii(s Receiver) {
	var art string
	if s.Allignflag == "right" {
		art = s.Allign(1)
	} else if s.Allignflag == "center" {
		art = s.Allign(2)
	} else if s.Allignflag == "justify" {
		art = s.Justify()
	} else {
		art = s.AllignLeft()
	}

	if IsFlagPassed("output") {
		validFileName, err := IsValidName(s.Outputflag)
		if !validFileName {
			fmt.Println(err)
			return
		}
		file, err := os.Create(s.Outputflag)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(art)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Print(art)
	}
	initial := Getwidth()
	if IsFlagPassed("align") {
		for {
			width := Getwidth()
			if width != initial {
				cmd := exec.Command("clear")
				stdout, err := cmd.Output()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(stdout))
				s.Spaces = nil
				s.Alligner()
				Ascii(s)
			}
		}
	}
}

// Space takes an int and returns a string with spaces of lenght n
func Space(n int) (space string) {
	for n > 0 {
		space += " "
		n--
	}
	return
}

// AllignLeft prints ascii-art the normal way
func (s *Receiver) AllignLeft() string {
	var count int
	reset := "\033[0m"
	var outputBuilder strings.Builder

	for _, val := range s.WordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(s.LettersToColor) == 0 && s.ColorCode != "" {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else if strings.Contains(s.LettersToColor, string(v)) {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else {
						outputBuilder.WriteString(s.FileArr[int(start)+i])
					}
				}
				outputBuilder.WriteString("\n")
			}
		} else {
			count++
			if count < len(s.WordsArr) {
				outputBuilder.WriteString("\n")
			}
		}
	}
	return outputBuilder.String()
}

// Allign prints the ascii either right or center
func (s *Receiver) Allign(n int) string {
	var count int
	reset := "\033[0m"
	var outputBuilder strings.Builder
	var art strings.Builder
	in := 0
	for _, val := range s.WordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(s.LettersToColor) == 0 && s.ColorCode != "" {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else if strings.Contains(s.LettersToColor, string(v)) {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else {
						outputBuilder.WriteString(s.FileArr[int(start)+i])
					}
				}
				outputBuilder.WriteString("\n")
				word := Space(s.Spaces[in]/n) + outputBuilder.String()
				art.WriteString(word)
				outputBuilder.Reset()
			}
		} else {
			count++
			if count < len(s.WordsArr) {
				art.WriteString("\n")
			}
		}
		in++
	}
	return art.String()
}

// Justify adds spaces in the middle to justify a string
func (s *Receiver) Justify() string {
	var count int
	reset := "\033[0m"
	var outputBuilder strings.Builder
	var art strings.Builder
	in := 0
	for _, val := range s.WordsArr {
		spaces := CountSpace(val)
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(s.LettersToColor) == 0 && s.ColorCode != "" {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else if strings.Contains(s.LettersToColor, string(v)) {
						outputBuilder.WriteString(s.ColorCode + s.FileArr[int(start)+i] + reset)
					} else {
						outputBuilder.WriteString(s.FileArr[int(start)+i])
					}
					if v == ' ' {
						outputBuilder.WriteString(Space(s.Spaces[in] / spaces))
					}
				}
				outputBuilder.WriteString("\n")
				word := outputBuilder.String()
				art.WriteString(word)
				outputBuilder.Reset()
			}
		} else {
			count++
			if count < len(s.WordsArr) {
				art.WriteString("\n")
			}
		}
		in++
	}
	return art.String()
}

// CountSpace takes a string and returns the number of spaces in the string
func CountSpace(word string) (count int) {
	for _, v := range word {
		if v == 32 {
			count++
		}
	}
	return
}
