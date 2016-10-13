package text

import "testing"
import "math"
import "fmt"

func TestSay(t *testing.T) {
	if Number(math.MaxInt64) != "nine quintillion two hundred and twenty three quadrillion three hundred and seventy two trillion thirty six billion eight hundred and fifty four million seven hundred and seventy five thousand eight hundred and seven" {
		t.Fail()
	}
	if Number(0) != "zero" {
		t.Fail()
	}
	if Number(-1000) != "minus one thousand" {
		t.Fail()
	}
	if Number(120) != "one hundred and twenty" {
		t.Fail()
	}
	if Number(2001) != "two thousand and one" {
		t.Fail()
	}
	fmt.Println(Price(150))
	//one fifty
	fmt.Println(Year(1980))
	//nineteen eighty
	fmt.Println(Year(2001))
	ThousandDivider = ","
	if Number(1123) != "one thousand, one hundred and twenty three" {
		t.Fail()
	}
}


