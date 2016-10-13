package text

import "testing"
import "math"
import "fmt"

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
	fmt.Println(Price(200))
}

/*  Hal3 Thu 13 Oct 02:22:39 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestSay
fifty pence
--- PASS: TestSay (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/number/text	0.009s
Thu 13 Oct 02:22:40 BST 2016
*/
/*  Hal3 Thu 13 Oct 02:22:49 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestSay
one pound, fifty
--- PASS: TestSay (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/number/text	0.004s
Thu 13 Oct 02:22:50 BST 2016
*/
/*  Hal3 Thu 13 Oct 02:23:03 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestSay
one pound, fifty one pence
--- PASS: TestSay (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/number/text	0.004s
Thu 13 Oct 02:23:04 BST 2016
*/
/*  Hal3 Thu 13 Oct 02:25:41 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestSay
one pound, fifty one pence
--- PASS: TestSay (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/number/text	0.003s
Thu 13 Oct 02:25:42 BST 2016
*/
/*  Hal3 Thu 13 Oct 02:32:50 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestSay
two pounds fifty five
--- PASS: TestSay (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/number/text	0.004s
Thu 13 Oct 02:32:51 BST 2016
*/
/*  Hal3 Thu 13 Oct 02:37:17 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestSay
two pounds
--- PASS: TestSay (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/number/text	0.003s
Thu 13 Oct 02:37:18 BST 2016
*/

