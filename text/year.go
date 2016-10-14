package text

import "strings"

// Year returns the textual representation of the given number in year format.
func Year(i int64) string {
	accumulator := make([]string, 0, 5)
	if i%1000 < 10 && i/1000 != 1 {
		number(i, &accumulator)
	} else {
		number(i/100, &accumulator)
		if i%100<10 {accumulator = append(accumulator, DateZero)}
		number(i%100, &accumulator)
	}
	return strings.Join(accumulator, Joiner)
}
