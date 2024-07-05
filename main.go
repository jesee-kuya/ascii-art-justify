package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	for _, v := range os.Args {
		if v == "--color" || v == "--output" || v == "--align" || v == "--align=" || v == "--output=" || v == "--color=" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		} else if strings.HasPrefix(v, "-color") || strings.HasPrefix(v, "-output") || strings.HasPrefix(v, "-align") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
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
		} else if receive.Allignflag == "right" {
			receive.AllignRight()
		} else if receive.Allignflag != "left" {
			fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
			return
		}
	}
	ascii.Ascii(receive)
}
