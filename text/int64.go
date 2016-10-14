package text

import "strings"

// Number converts a number into text, in English speech.
func Number(i int64) string {
	if i == 0 {
		return Units[0]
	}
	accumulator := make([]string, 0, 5)
	if i < 0 {
		if len(Negative) > 0 {
			accumulator = append(accumulator, Negative)
		}
		i = -i
	} else {
		if len(Positive) > 0 {
			accumulator = append(accumulator, Positive)
		}
	}
	number(i, &accumulator)
	return strings.Join(accumulator, Joiner)
}

func number(n int64, accumulator *[]string) {
	i:=n
	if i >= 1e18 {
		lessThan1000(i/1e18, accumulator)
		i %= 1e18
		*accumulator = append(*accumulator, PowerThousands[5]+ThousandDivider)
	}
	if i >= 1e15 {
		lessThan1000(i/1e15, accumulator)
		i %= 1e15
		*accumulator = append(*accumulator, PowerThousands[4]+ThousandDivider)
	}
	if i >= 1e12 {
		lessThan1000(i/1e12, accumulator)
		i %= 1e12
		*accumulator = append(*accumulator, PowerThousands[3]+ThousandDivider)
	}
	if i >= 1e9 {
		lessThan1000(i/1e9, accumulator)
		i %= 1e9
		*accumulator = append(*accumulator, PowerThousands[2]+ThousandDivider)
	}
	if i >= 1e6 {
		lessThan1000(i/1e6, accumulator)
		i %= 1e6
		*accumulator = append(*accumulator, PowerThousands[1]+ThousandDivider)
	}

	if i >= 1e3 {
		lessThan1000(i/1e3, accumulator)
		i %= 1e3
		*accumulator = append(*accumulator, PowerThousands[0]+ThousandDivider)
	}
	if i>0 && i!=n && i/100==0 && And!="" { *accumulator = append(*accumulator, And)}
	lessThan1000(i, accumulator)
}

func lessThan100(i int64, accumulator *[]string) {
	if i < 20 {
		*accumulator = append(*accumulator, Units[i])
	} else {
		*accumulator = append(*accumulator, Tens[i/10-2])
		if i%10 != 0 {
			*accumulator = append(*accumulator, Units[i%10])
		}
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
			if And !=""  {*accumulator = append(*accumulator, And)}
		}
		lessThan100(i%100, accumulator)
	}
}


