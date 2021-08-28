package main

import "fmt"

func main() {
	// make array of 3 integers
	var a [3]int
	fmt.Println("elemen array indeks ke 0: ", a[0])
	fmt.Println("panjang array a: ", len(a))

	// print the indices and the element
	for i, v := range a {
		fmt.Printf("indeks ke-%d nilainya %d\n", i, v)
	}

	// print the element only
	for _, v := range a {
		fmt.Printf("nilai di dalam array %d\n ", v)
	}

	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n ", q)
}