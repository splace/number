package text

import "testing"
import "math"

func TestSay(t *testing.T) {
	if Number(math.MaxInt64) != "nine quintillion two hundred and twenty three quadrillion three hundred and seventy two trillion thirty six billion eight hundred and fifty four million seven hundred and seventy five thousand eight hundred and seven" {
		t.FailNow()
	}
	if Number(0) != "zero" {
		t.FailNow()
	}
	if Number(-1000) != "minus one thousand" {
		t.FailNow()
	}
	if Number(123) != "one hundred and twenty three" {
		t.FailNow()
	}
	ThousandDivider="," 
	if Number(1123) != "one thousand, one hundred and twenty three" {
		t.FailNow()
	}
}

