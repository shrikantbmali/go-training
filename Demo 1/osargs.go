package main

import (
	"fmt"
	"os"
	"strconv"

	"../src/calc"
	"../src/thrid"
)

func main() {
	fmt.Println("My favorite number is", os.Args)

	a, _ := strconv.ParseInt(os.Args[1], 10, 0)
	b, _ := strconv.ParseInt(os.Args[2], 10, 0)

	fmt.Println("Sum of 30, and 70 is ", calc.Add(a, b))
	fmt.Println("Devision of 70, and 30 is ", calc.Devide(b, a))
	fmt.Println("Multiplication of 70, and 30 is ", thrid.Mult(b, a))
}
