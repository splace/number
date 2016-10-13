package text

import "math"
import "fmt"

func ExampleMaxInt() {
	fmt.Println(Number(math.MaxInt64))
	/* Output:
	nine quintillion two hundred and twenty three quadrillion three hundred and seventy two trillion thirty six billion eight hundred and fifty four million seven hundred and seventy five thousand eight hundred and seven
	*/
}

func Example0() {
	fmt.Println(Number(0))
	/* Output:
	zero
	*/
}
func Example1000() {
	fmt.Println(Number(-1000))
	/* Output:
	minus one thousand
	*/
}
func Example120() {
	fmt.Println(Number(120))
	/* Output:
	one hundred and twenty
	*/
}
func Example3011() {
	fmt.Println(Number(3011))
	/* Output:
	three thousand and eleven
	*/
}
func Example3000001() {
	fmt.Println(Number(3000001))
	/* Output:
	three million and one
	*/
}
func Example3000000001() {
	fmt.Println(Number(3000000001))
	/* Output:
	three billion and one
	*/
}
func Example150() {
	fmt.Println(Price(150))
	/* Output:
	one fifty
	*/
}
func Example1980() {
	fmt.Println(Year(1980))
	/* Output:
	nineteen eighty
	*/
}
func Example2001() {
	fmt.Println(Year(2001))
	/* Output:
	two thousand and one
	*/
}	
func Example1123() {
	ThousandDivider = ","
	fmt.Println(Number(1123))
	/* Output:
	one thousand, one hundred and twenty three
	*/
}


