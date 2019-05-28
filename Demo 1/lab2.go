package main

import "fmt"

func main() {
	var c1, c2 = "Hello", "World"
	var a, b, c = 0, 1.23, false
	x := 0
	y := 1.23
	z := false

	fmt.Printf("c1 %s, C2 %s\n", c1, c2)
	fmt.Printf("a %d, b %f, c %t\n", a, b, c)
	fmt.Printf("x %d, y %f, z %t\n", x, y, z)

	fmt.Println("c1 ", c1, "C2 ", c2)
	fmt.Println("a ", a, "b ", b, "c ", c)
	fmt.Println("x ", x, "y ", y, "z ", z)
}
