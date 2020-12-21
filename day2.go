package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var y float64
	var z float64
	var aa float64
	var a float64

	fmt.Print("Enter the value of x \n")
	fmt.Scan(&x)
	fmt.Print("Enter the value of y \n")
	fmt.Scan(&y)
	fmt.Print("Enter the value of z \n")
	fmt.Scan(&z)
	fmt.Print("Enter the value of aa \n")
	fmt.Scan(&aa)

	a = ((math.Sqrt(x) + math.Sqrt(y)*z) / aa)

	fmt.Println("result of (root(x) + root(y)*z)/aa) = ", a)

}
