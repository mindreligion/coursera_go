package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a float number")
	var x float64
	fmt.Scan(&x)
	fmt.Println(int(x))
}
