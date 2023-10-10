package main

import (
	"fmt"
	"strconv"
)

//	func greeting(s string) string {
//		return "hello " + s
//	}
func main() {
	//var n int
	// fmt.Println("password")
	// //fmt.Scanln(&n)

	// fmt.Println(greeting("Abhi"))
	_, err := Addition("1o", "10")
	if err != nil {
		fmt.Println(err)
		return
	}
	//	fmt.Println(e)
	//fmt.Println(Addition("10", "2"))

}
func Addition(a, b string) (int, error) {
	x, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return x + y, nil
}
