package text

import "strings"


// Number returns the string representation (in English) of the number provided.
func Number(i int64) string {
	if i==0 {return Units[0]}
	accumulator := make([]string, 0, 5)
	if i<0 {
		if len(Negative)>0{accumulator = append(accumulator, Negative)}
		i=-i
	}else{
		if len(Positive)>0{accumulator = append(accumulator, Positive)}
	}
	number(i, &accumulator)
	return strings.Join(accumulator, Joiner)
}

func number(i int64, accumulator *[]string) {
	if i >= 1e18 {
		lessThan1000(i/1e18, accumulator)
		*accumulator = append(*accumulator, PowerThousands[5]+ThousandDivider)
		i %= 1e18
	}
	if i >= 1e15 {
		lessThan1000(i/1e15, accumulator)
		*accumulator = append(*accumulator, PowerThousands[4]+ThousandDivider)
		i %= 1e15
	}
	if i >= 1e12 {
		lessThan1000(i/1e12, accumulator)
		*accumulator = append(*accumulator, PowerThousands[3]+ThousandDivider)
		i %= 1e12
	}
	if i >= 1e9 {
		lessThan1000(i/1e9, accumulator)
		*accumulator = append(*accumulator, PowerThousands[2]+ThousandDivider)
		i %= 1e9
	}
	if i >= 1e6 {
		lessThan1000(i/1e6, accumulator)
		*accumulator = append(*accumulator, PowerThousands[1]+ThousandDivider)
		i %= 1e6
	}

	if i >= 1e3 {
		lessThan1000(i/1e3, accumulator)
		*accumulator = append(*accumulator, PowerThousands[0]+ThousandDivider)
		i %= 1e3
	}
	lessThan1000(i, accumulator)
}

func lessThan100(i int64, accumulator *[]string) {
	if i == 0 {
		return
	}
	if i < 20 {
		*accumulator = append(*accumulator, Units[i])
	} else {
		*accumulator = append(*accumulator, Tens[i/10-2])
		*accumulator = append(*accumulator, Units[i%10])
	}
}

func lessThan1000(i int64, accumulator *[]string) {
	if i == 0 {
		return
	}
	if i < 100 {
		lessThan100(i, accumulator)
	} else {
		*accumulator = append(*accumulator, Units[i/100])
		*accumulator = append(*accumulator, Hundred)
		if i%100 > 0 {
			*accumulator = append(*accumulator, HundredDivider)
		}
		lessThan100(i%100, accumulator)
	}
}


