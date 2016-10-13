package text

import "strings"


// Price returns the textual representation of the given number (small currency units) in English.
// You will need to multiply a main currency unit by 100.
func Price(i int64) string {
	if i==0 {return CurrencyZero}
	accumulator := make([]string, 0, 10)
	if i<0 {
		if len(Negative)>0{accumulator = append(accumulator, Negative)}
		i=-i
	}else{
		if len(Positive)>0{accumulator = append(accumulator, Positive)}
	}
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

