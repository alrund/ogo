package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const (
	slash int = iota + 1
	letter
	digit
)

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	err := checkString(str)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	for i := 0; i < len(str); i++ {
		err = add(&result, []rune(str), i)
		if err != nil {
			return "", err
		}
	}
	return result.String(), err
}

func add(result *strings.Builder, str []rune, i int) error {
	r := str[i]
	rt := runeType(r)
	rpt, rppt, rpppt, rnt, rn := nearestValues(str, i)

	switch rt {
	case letter:
		return addLetter(result, rnt, r, rn)

	case slash:
		return addSlash(result, rpt, rppt, rnt, r, rn)

	case digit:
		return addDigit(result, rpt, rppt, rpppt, rnt, r, rn)
	}

	return nil
}

func addDigit(result *strings.Builder, rpt, rppt, rpppt, rnt int, r, rn rune) error {
	// \45
	if rnt == digit && rpt == slash {
		err := addRepeatedRune(result, r, rn)
		if err != nil {
			return err
		}
		return nil
	}
	// \3 | \\\3
	if rpt == slash && (rppt != slash || rpppt == slash) {
		result.WriteString(string(r))
		return nil
	}

	return nil
}

func addSlash(result *strings.Builder, rpt, rppt, rnt int, r, rn rune) error {
	// \5
	if rpt == slash && rnt == digit && rppt != slash {
		err := addRepeatedRune(result, r, rn)
		if err != nil {
			return err
		}
		return nil
	}

	// \\
	if rpt == slash && rppt != slash {
		result.WriteString(string(r))
		return nil
	}

	return nil
}

func addLetter(result *strings.Builder, rnt int, r, rn rune) error {
	// a4
	if rnt == digit {
		err := addRepeatedRune(result, r, rn)
		if err != nil {
			return err
		}
		return nil
	}
	// ab | e\ | last
	if rnt == letter || rnt == slash || rnt == 0 {
		result.WriteString(string(r))
		return nil
	}

	return nil
}

func nearestValues(str []rune, i int) (int, int, int, int, rune) {
	var (
		rn                    rune
		rpt, rppt, rpppt, rnt int
	)

	if i-3 >= 0 {
		rpppt = runeType(str[i-3])
	}
	if i-2 >= 0 {
		rppt = runeType(str[i-2])
	}
	if i-1 >= 0 {
		rpt = runeType(str[i-1])
	}
	if i+1 < len(str) {
		rn = str[i+1]
		rnt = runeType(rn)
	}

	return rpt, rppt, rpppt, rnt, rn
}

func addRepeatedRune(result *strings.Builder, r, n rune) error {
	num, err := strconv.Atoi(string(n))
	if err != nil {
		return err
	}
	result.WriteString(strings.Repeat(string(r), num))

	return nil
}

func runeType(r rune) int {
	if r == '\\' {
		return slash
	}

	if unicode.IsDigit(r) {
		return digit
	}

	return letter
}

func checkString(str string) error {
	if unicode.IsDigit(rune(str[0])) {
		return ErrInvalidString
	}

	var (
		matched bool
		err     error
	)

	matched, err = regexp.Match(`[^\\]\d{2}`, []byte(str))
	if err != nil {
		return err
	}

	if matched {
		return ErrInvalidString
	}

	matched, err = regexp.Match(`\\[^\d\\]`, []byte(str))
	if err != nil {
		return err
	}

	if matched {
		return ErrInvalidString
	}

	return nil
}
