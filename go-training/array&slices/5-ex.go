package main

import "fmt"

func main() {
	// a := make([]int, 5, 5) //make is use to create the slices with predefining the capacity
	// fmt.Println(len(a))
	// fmt.Println(cap(a))
	// fmt.Println(a)
	// fmt.Printf("%p", a)
	// a = append(a, 100, 2, 3, 4, 5, 6)
	// fmt.Println()
	// fmt.Printf("%p", a)
	a := []int{10, 20, 30, 40, 50}
	b := make([]int, len(a[1:3]))
	copy(b, a[1:3])
	// fmt.Println(a)
	// fmt.Println(b)
	//b[0] = 100
	//b = append(b, 22, 33)
	fmt.Println(a)
	fmt.Println(b)
}
