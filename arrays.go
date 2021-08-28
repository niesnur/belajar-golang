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

	fmt.Println("---------------------------")
	// make empty array of 5 integers
	var x [5] int
	fmt.Println("emp: ", x)

	// set value of x[4] 
	x[4] = 100
	fmt.Println("set: ", x)
	fmt.Println("get value of x[4]: ", x[4]) 

	// print length of x
	fmt.Println("len: ", len(x))

	y := [5] int {1, 2, 3, 4, 5}
	fmt.Println("dcl: ", y)

	var twoD [2][3] int
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}