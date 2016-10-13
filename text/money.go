package text

import "strings"


// Price returns the textual representation as money (in English) of the number provided divided by 100.
// (so you will need to multiply a main unit by 100)
func Price(i int64) string {
	if i==0 {return CurrencyZero}
	accumulator := make([]string, 0, 10)
	number(i/100, &accumulator)
	switch i/100 {
	case 0:
		number(i%100, &accumulator)
		accumulator = append(accumulator, SmallCurrency)
	case 1:
		if i%10>0 {accumulator = append(accumulator, Currency)}
		number(i%100, &accumulator)
	case 2,3,4,5,6,7,8,9:
		accumulator = append(accumulator, Currency+"s")
		number(i%100, &accumulator)
	default:
		accumulator = append(accumulator, Currency+"s"+SmallDivider)
		number(i%100, &accumulator)
	}
	return strings.Join(accumulator, Joiner)
}


