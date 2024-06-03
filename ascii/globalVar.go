package ascii

var Colormap = map[string]string{
	"yellow":  "\033[33;1m",
	"red":     "\033[31;1m",
	"green":   "\033[32;1m",
	"blue":    "\033[34;1m",
	"magenta": "\033[35;1m",
	"cyan":    "\033[36;1m",
	"gray":    "\033[37;1m",
	"white":   "\033[97;1m",
	"black":   "\033[30;1m",
	"purple":  "\033[38;2;128;0;128m",
	"orange":  "\033[38;2;255;165;0m",
	"brown":   "\033[38;2;165;42;42m",
	"pink":    "\033[38;2;255;192;203m",
	"gold":    "\033[38;2;255;215;0m",
	"silver":  "\033[38;2;192;192;192m",
	"violet":  "\033[38;2;238;130;238m",
	"maroon":  "\033[38;2;128;0;0m",
	"navy":    "\033[38;2;0;0;128m",
	"olive":   "\033[38;2;128;128;0m",
}

type RGB struct {
	R int
	G int
	B int
}
