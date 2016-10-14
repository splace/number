package text

import "math"
import "fmt"

func ExampleNumber() {
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
func ExamplePrice() {
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
func ExampleYear() {
	fmt.Println(Year(-5001))
	/* Output:
	five thousand and one BC
	*/
}	
func Example1904() {
	fmt.Println(Year(1904))
	/* Output:
	nineteen oh four
	*/
}
func Example25() {
	fmt.Println(Year(25))
	/* Output:
	twenty five
	*/
}
func Example1123() {
	ThousandDivider = ","
	fmt.Println(Number(1123))
	/* Output:
	one thousand, one hundred and twenty three
	*/
}

func Example1630() {
	fmt.Println(TimeOfDay(16*60+30))
	/* Output:
	four thirty
	*/
}

func Example2105() {
	fmt.Println(TimeOfDay(21*60+5))
	/* Output:
	nine oh five
	*/
}

func Example230() {
	fmt.Println(TimeOfDay(2*60+10))
	/* Output:
	two ten am
	*/
}


func ExampleTimeOfDay() {
	fmt.Println(TimeOfDay(16*60))
	/* Output:
	four o'clock
	*/
}


func ExampleMidnight() {
	fmt.Println(TimeOfDay(0))
	/* Output:
	12 midnight
	*/
}



