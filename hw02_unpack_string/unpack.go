package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	err := validateString(str)
	if err != nil {
		return "", err
	}
	runes := []rune(str)
	length := len(runes)
	var result strings.Builder
	for i := 0; i < length; {
		var d int
		var r, rn, rnn rune
		r = runes[i]
		if i+1 < length {
			rn = runes[i+1]
		}
		if i+2 < length {
			rnn = runes[i+2]
		}
		switch {
		case isLetter(r) && isDigit(rn):
			err = addRepeatedRune(&result, r, rn)
			d = 2
		case isLetter(r) && (isLetter(rn) || isSlash(rn)):
			result.WriteString(string(r))
			d = 1
		case isSlash(r) && isDigit(rn) && !isDigit(rnn):
			result.WriteString(string(rn))
			d = 2
		case isSlash(r) && isDigit(rn) && isDigit(rnn):
			err = addRepeatedRune(&result, rn, rnn)
			d = 3
		case isSlash(r) && isSlash(rn) && isDigit(rnn):
			err = addRepeatedRune(&result, rn, rnn)
			d = 3
		case isSlash(r) && isSlash(rn):
			result.WriteString(string(rn))
			d = 2
		}
		if err != nil {
			return "", err
		}
		i += d
	}
	return result.String(), err
}

func addRepeatedRune(result *strings.Builder, r, n rune) error {
	num, err := strconv.Atoi(string(n))
	if err != nil {
		return err
	}
	result.WriteString(strings.Repeat(string(r), num))
	return nil
}

func isSlash(r rune) bool {
	return r == '\\'
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isLetter(r rune) bool {
	return r != '\\' && !unicode.IsDigit(r)
}

func validateString(str string) error {
	if unicode.IsDigit(rune(str[0])) {
		return ErrInvalidString
	}
	rules := []string{`[^\\]\d{2}`, `\\[^\d\\]`}
	for _, rule := range rules {
		r, err := regexp.Compile(rule)
		if err != nil {
			return err
		}
		if matched := r.Match([]byte(str)); matched {
			return ErrInvalidString
		}
	}
	return nil
}
