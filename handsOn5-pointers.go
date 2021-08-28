package main

import "fmt"

func main() {
	var i, j = 42, 2701
	var p = &i

	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37

	fmt.Println(j)
}