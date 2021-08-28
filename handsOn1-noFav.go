package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("Nomor Favorit saya adalah: ", rand.Intn(10))
}