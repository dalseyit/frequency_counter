package bytes

import (
	"bytes"
	"unicode"
)

// SplitIntoWords splits given slice into words.
// A continuous unicode letters will be considered as a word
func SplitIntoWords(b []byte) [][]byte {
	fieldsFunc := func(r rune) bool { return !unicode.IsLetter(r) }
	return bytes.FieldsFunc(b, fieldsFunc)
}
