package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println(b)
	c := [...]int{10, 20, 50, 50, 60, 60, 60}
	fmt.Println(len(c))
}
