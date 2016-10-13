package text

import "strings"


// Year returns the textual representation of the given number in yeay format English.
func Year(i int64) string {
	number(i/100, &accumulator)
	number(i%100, &accumulator)
	return strings.Join(accumulator, Joiner)
}


