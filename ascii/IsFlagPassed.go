package ascii

import "flag"

// IsFlagPassed checks if a flag with the given name is passed in the command line
func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
