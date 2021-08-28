package main

import (
	"fmt"
	"math"
)

func main() {
	// membuat var x dan y dengan tipe data int bernilai 3 & 4
	var x, y int = 3, 4

	// menghitung akar (x*x + y*y) dengan math.Sqrt
	// membuat var z dengan tipe data unint
	z := uint(math.Sqrt(float64(x*x + y*y)))
	fmt.Println(x, y, z)
}