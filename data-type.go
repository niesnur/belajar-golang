package main

import "fmt"

func main() {
	// string
	fmt.Println("go" + "lang")
	// integer
	fmt.Println("1 + 1 =", 1+1)
	//float
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// decimal
	var decimalNumber float64 = 2.62
	fmt.Printf("Decimal Number: %f\n", decimalNumber)
	fmt.Printf("Decimal Number: %.3f\n", decimalNumber)

	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	s = `hello, world
		 mari belajar golang dengan kami`
	fmt.Println(s)

	const pi = 3.14159
	const constant string = "constant"
	fmt.Println(pi, constant)
}