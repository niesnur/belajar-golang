package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Printf("\n")

	a := 1
	for a < 100{
		fmt.Print(a, " ")
		a = a * 2
	}

	fmt.Printf("\n")
	
	evenVals := [] int {2, 4, 6, 8, 10, 12}
	for x, v := range evenVals{
		fmt.Println(x, v)
	}
}