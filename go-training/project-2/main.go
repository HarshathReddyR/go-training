package main

import (
	"fmt"
	"project-2/db"
)

func main() {
	// a, b := db.New("abcd")
	// fmt.Println(a, b)
	c1 := db.Conn{
		db: "hrr",
	}
	fmt.Println(c1)
	a, b := db.New("abcd")
	fmt.Println(a, b)
}
