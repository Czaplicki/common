package strutil

import "errors"

func GetSuffix(sequence string, suffixdivision rune) (string, error) {
	data := []rune(sequence)
	for i := len(data) -1; i > -1; i-- {
		if data[i] == suffixdivision {
			return string(data[len(data) - 1:]), nil
		}
	}
	return "", errors.New("suffix not found on sequence : " + sequence)
}
