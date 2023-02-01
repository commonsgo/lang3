package stringutils

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const PAD_LIMIT = 8192
const EMPTY = ""

func Capitalize(s string) string {

	c1 := rune(s[0])
	capc1 := unicode.ToTitle(c1)
	var b strings.Builder

	b.Grow(len(s))

	b.WriteRune(capc1)
	b.WriteString(s[1:])

	return b.String()
}

func Center(s string, size int) string {

	if size <= 0 {
		return s
	}
	strLen := len(s)
	pads := size - strLen

	if pads <= 0 {
		return s
	}
	s1 := LeftPad2(s, strLen+pads/2, ' ')
	s1 = RightPad(s1, size, ' ')

	return s1
}

func LeftPad2(s string, pads int, padChar rune) string {
	// deal with null
	if pads <= 0 {
		return s
	}
	netPads := pads - len(s)
	if netPads <= 0 {
		return s
	}

	return Repeat(padChar, netPads) + s
}

func RightPad(s string, pads int, padChar rune) string {

	return ""
}

func LeftPad(s string, targetSize int, padStr string) string {

	if len(padStr) == 0 {
		padStr = " "
	}

	strLen := len(s)

	pads := targetSize - strLen

	if pads <= 0 {
		return s
	}

	padStrLen := len(padStr)

	if padStrLen == 1 && pads <= PAD_LIMIT {

		runeValue, _ := utf8.DecodeRuneInString(padStr)
		return LeftPad2(s, targetSize, runeValue)

	}

	if pads == padStrLen {
		return padStr + s
	} else if pads < padStrLen {
		return padStr + s
		// how to substring?
	} else {
		var b strings.Builder
		b.Grow(targetSize)
		for i := 0; i < pads; i++ {
			b.WriteString(padStr)
		}
		b.WriteString(s)
		return b.String()
	}

}

func IsEmpty(s string) bool {

	if len(s) == 0 {
		return true
	}

	return false
}

func Repeat(c rune, repeat int) string {

	if repeat <= 0 {
		return EMPTY
	}
	var b strings.Builder
	b.Grow(repeat)
	for i := 0; i < repeat; i++ {
		b.WriteRune(c)
	}
	return b.String()
}

func Wrap(s string, c rune) string {
	//if c == nil or empty char or unprintable char?

	// if s is nil?

	var b strings.Builder
	b.Grow(len(s) + 2)
	b.WriteRune(c)
	b.WriteString(s)
	b.WriteRune(c)
	return b.String()
}

func WrapStr(s string, wrapper string) string {
	//if c == nil or empty char or unprintable char?

	// if s is nil?

	var b strings.Builder
	b.Grow(len(s) + 2)
	b.WriteString(wrapper)
	b.WriteString(s)
	b.WriteString(wrapper)
	return b.String()
}
