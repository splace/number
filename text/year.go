package text

import "strings"

// Year converts a year, into text in English speech.
func Year(i int64) string {
	accumulator := make([]string, 0, 5)
	ending:=DatePositive
	if i<0{i=-i;ending=DateNegative}
	if i%1000 < 10 && i/1000 != 1 {
		number(i, &accumulator)
	} else {
		number(i/100, &accumulator)
		if i%100<10 {accumulator = append(accumulator, DateZero)}
		number(i%100, &accumulator)
	}
	if ending!="" {accumulator = append(accumulator, ending)}
	return strings.Join(accumulator, Joiner)
}
