package main

import (
	"flag"
	"fmt"
	"os"

	"ascii/ascii"
)

// main function is used to read the banner file and print the ascii art
// based on the arguments passed
func main() {
	for _, v := range os.Args {
		if v == "--color" || v == "-color" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		} else if v == "--output" || v == "-output" {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		}
	}
	var filename string
	var colorflag string
	var lettersTocolor string
	var colorCode string
	var outputFlag string

	flag.StringVar(&outputFlag, "output", "banner.txt", "specify the ascii-font to be printed")
	flag.StringVar(&colorflag, "color", "reset", "color for color input")
	flag.Parse()
	argsPassed := flag.Args()

	if len(argsPassed) == 3 {
		if !ascii.IsFlagPassed("color") {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		}
		filename = argsPassed[2]
		argsPassed = argsPassed[:2]
	} else if len(argsPassed) == 2 {
		if ascii.IsFlagPassed("color") {
			_, err := ascii.GetFileName(argsPassed[1])
			if err != nil {
				filename = "standard"
			} else {
				filename = argsPassed[1]
				argsPassed = argsPassed[:1]
			}
		} else {
			filename = argsPassed[1]
			argsPassed = argsPassed[:1]
		}
	} else if len(argsPassed) == 1 {
		filename = "standard"
	}

	bannerContent, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ascii.IsFlagPassed("color") {
		ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent, outputFlag)
	} else {
		colorCode = ""
		ascii.Art(argsPassed, bannerContent, lettersTocolor, colorCode, 0, outputFlag)
	}
}
