package resource_prefix

import (
	"math/bits"
	"strings"
)

// Go strings are UTF-8, and these characters all map to single bytes, so this is like a `const`
// slice of bytes for the possible characters (a normal slice can't be a constant).
//
//goland:noinspection SpellCheckingInspection
const characterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

// The domain of what can be represented by GenerateResourcePrefix
const possibleCombinations = uint64(26 * 36 * 36 * 36)

// GenerateResourcePrefix converts the input number to four characters, written into the string builder.
//
// The string is like `[r3][r2][r1][r0]`.
// r0 through r2 are in base 36, for 26 letters plus 10 digits.
// r3 is in base 26, for 26 letters, so the string always starts with a letter.
// r0 is the "lowest" digit, and the string is a bit similar to a hexadecimal number with letters taking
// on numeral values. The result is a string representation of the input number, modulo the domain of
// the resulting string, achieving full coverage of that domain to minimize conflicts.
//
// Why does this function exist? For ephemeral environments it is sometimes required to have a short
// differentiating string that can be a prefix onto cloud resource names to try to avoid conflicts.
// Four characters is the largest this string can be, so we just do our best. The idea is that
// Sherlock can generate a prefix from the count of all environments, check if it already exists,
// and keep incrementing the number until it finds an opening. Since ephemeral environments are
// deleted with some regularity, we can reasonably assume this won't grow out of control.
//
// For more historical context, environments imported to Sherlock actually included resource prefixes
// already, so the algorithm needs to be able to handle surprise conflicts (we can't just go based on
// environment ID or something).
//
// Examples (remember that input is always modulo possibleCombinations):
// possibleCombinations-2 => z998
// possibleCombinations-1 => z999
// possibleCombinations   => aaaa
// possibleCombinations+1 => aaab
// possibleCombinations+2 => aaac
func GenerateResourcePrefix(sb *strings.Builder, number uint64) {
	number, r0 := bits.Div64(0, number%possibleCombinations, 36)
	number, r1 := bits.Div64(0, number, 36)
	number, r2 := bits.Div64(0, number, 36)
	_, r3 := bits.Div64(0, number, 26)
	sb.WriteByte(characterBytes[r3])
	sb.WriteByte(characterBytes[r2])
	sb.WriteByte(characterBytes[r1])
	sb.WriteByte(characterBytes[r0])
}
