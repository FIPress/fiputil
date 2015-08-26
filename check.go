package fiputil

import (
	"bytes"
	"fmt"
	"strconv"
)

func IsSpace(c byte) bool {
	switch c {
	case ' ', '\t':
		return true
	}
	return false
}

func IsLineEnd(c byte) bool {
	switch c {
	case '\n', '\r', '\f':
		return true
	}
	return false
}

//A blank line is any line that looks like a blank line
// — a line containing nothing but spaces or tabs is considered blank
func IsBlankLine(input []byte) (bool, int) {
	i := 0
L:
	for ; i < len(input); i++ {
		switch input[i] {
		case '\n', '\r', '\f':
			return true, i + 1
		case ' ', '\t':
			continue L
		default:
			return false, 0
		}
	}
	return true, i
}

//A block end by one or more blank lines.
func IsBlockEnd(input []byte) bool {
	if !IsLineEnd(input[0]) {
		return false
	}
	isBlank, _ := IsBlankLine(input[1:])
	return isBlank
}

func IsSpaceOrLineEnd(c byte) bool {
	switch c {
	case ' ', '\t', '\n', '\r', '\f':
		return true
	}
	return false
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsIdentifier(c byte) bool {
	return IsDigit(c) || IsLetter(c)
}

func IsIndented(input []byte, indentSize int) bool {
	if len(input) == 0 {
		return false
	}
	if input[0] == '\t' {
		return true
	}

	if IsSpaceIndented(input, indentSize) {
		return true
	}

	return false
}

func IsSpaceIndented(input []byte, indentSize int) bool {
	if len(input) < indentSize {
		return false
	}
	for i := 0; i < indentSize; i++ {
		if input[i] != ' ' {
			return false
		}
	}
	return true
}

func SkipSpace(input []byte) int {
	i := 0
	for i < len(input) {
		if IsSpace(input[i]) {
			i++
		} else {
			return i
		}
	}
	return i
}

func SkipSpaceAndLineEnd(input []byte) int {
	i := 0
	for i < len(input) {
		if IsSpaceOrLineEnd(input[i]) {
			i++
		} else {
			return i
		}
	}
	return i
}

//Skip until char
func SkipUntil(input []byte, c byte) (idx int, found bool) {
	i := 0
	for ; i < len(input); i++ {
		if input[i] == c {
			return i, true
		}
	}
	return i, false
}

/*
return the chars skipped. -1 indicates there is no match
*/
func SkipUntilArray(input []byte, arr []byte) (delta int, found bool) {
	length := len(input)
	arrLength := len(arr)
	for i := 0; i < length; i++ {
		if input[i] == arr[0] {
			if i+arrLength <= length &&
				bytes.Compare(input[i+1:i+arrLength], arr[1:]) == 0 {
				return i, true
			}
		}
	}
	return 0, false
}

func SkipUntilFunc(input []byte, f func(byte) bool, include bool) int {
	i := 0

	for i < len(input) {
		if f(input[i]) {
			if include {
				return i + 1
			}
			return i
		} else {
			i++
		}
	}
	return i
}

func SkipUntilOrStopAtLineEnd(input []byte, until byte) (idx int, found bool) {
	i := 0
	for ; i < len(input); i++ {
		if input[i] == until {
			return i, true
		} else if IsLineEnd(input[i]) {
			return i, false
		}
	}
	return i, false
}

func SkipUntilOrStopAt(input []byte, until byte, stop byte) (idx int, found bool) {
	i := 0
	for ; i < len(input); i++ {
		if input[i] == until {
			return i, true
		} else if input[i] == stop {
			return i, false
		}
	}
	return i, false
}

func SkipUntilOrStop(input []byte, until byte, stopFunc func([]byte) bool) (idx int, found bool) {
	i := 0
	for i < len(input) {
		if input[i] == until {
			return i + 1, true
		} else if stopFunc(input[i:]) {
			return i, false
		} else {
			i++
		}
	}
	return i, false
}

//escape
func Unquote(in string) string {
	if len(in) == 0 {
		return in
	}

	out := fmt.Sprintf("%c%s%c", '"', in, '"')
	out, err := strconv.Unquote(out)
	if err != nil {
		return in
	}
	return out
}
