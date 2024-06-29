package main

import (
	"flag"
	"fmt"
	"os"

	"ascii/ascii"
)

func main() {
	for _, v := range os.Args {
		if v == "--color" || v == "-color" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	var receive ascii.Receiver

	flag.StringVar(&receive.Colorflag, "color", "reset", "color for color input")
	flag.StringVar(&receive.Outputflag, "output", "banner.txt", "Write ascii output to a specified file")
	flag.StringVar(&receive.Allignflag, "align", "justify", "Set the alignment of ascii art on the terminal")
	flag.Parse()
	argsPassed := flag.Args()

	msg := receive.SortArg(argsPassed)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	if ascii.IsFlagPassed("color") {
		receive.Color()
	} else {
		receive.ColorCode = ""
	}
	receive.Art()
	if ascii.IsFlagPassed("align") {
		if receive.Allignflag == "center" {
			receive.AllignCentre()
		} else if receive.Allignflag == "justify" {
			receive.AllignJustify()
		}
	}
	ascii.Ascii(receive)
}
