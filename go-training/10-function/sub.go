package main

import "fmt"

func main() {
	// operate(add)
	// operate(sub)
	// operate(mul)
	operate(add, 10, 50)
}

// func operate(op func(x, y int) int) {
// 	fmt.Println(op(10, 20))
// }

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}

func operate(op func(x, y int) int, a, b int) {
	fmt.Println(op(10, 20))
}

// type operate func(a, b int)
