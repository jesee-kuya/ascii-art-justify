package ascii

import (
	"errors"
	"strings"
)

func ValidCharacter(r rune) bool {
	specialCh := "^&'@{}[],$=!-#()%.+~_"
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	for _, v := range specialCh {
		if v == r {
			return true
		}
	}
	return false
}

func IsValidName(str string) (bool, error) {
	if !strings.HasSuffix(str, ".txt") {
		return false, errors.New("wrong file name")
	}

	for i, v := range str {
		if i == 0 && (v == ' ' || v == '.' || v == '-' || v == '_') {
			return false, errors.New("wrong file name")
		} else if i == (len(str)-5) && (v == ' ' || v == '.' || v == '-' || v == '_') {
			return false, errors.New("wrong file name")

		} else if !ValidCharacter(v) {
			return false, errors.New("wrong file name")
		}
	}
	return true, nil
}
