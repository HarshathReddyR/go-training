package main

import "fmt"

func main() {
	// var i []int
	// fmt.Printf("%v", i)
	// var i []int
	// i[1] = 100 // update
	// fmt.Println(i)

	var i []int
	//i[1] = 100 // update
	//fmt.Println(i)
	if i == nil {
		fmt.Println("slice is nil", i)
	}
	i = append(i, 11, 12)
	fmt.Println(i)
	b := []int{10, 230, 4050, 60}
	fmt.Println(b)
}
