package reverse

import (
	"unicode/utf8"
)

// reverse reverses the characters of a UTF-8 encoded []byte slice in place.
func reverse(b []byte) {
	// Step 1: Reverse the byte slice entirely
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	// Step 2: Reverse each individual UTF-8 character to correct rune order
	for i := 0; i < len(b); {
		// Decode the reversed rune
		r, size := utf8.DecodeRune(b[i:])
		// Reverse the bytes of the current rune in place
		reverseBytes(b[i : i+size])
		i += size
	}
}

// reverseBytes reverses a slice of bytes in place
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
