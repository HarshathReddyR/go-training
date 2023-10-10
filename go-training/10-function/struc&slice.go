package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	u := []user{
		{
			name: "Abhi",
			age:  30,
		},
		{
			name: "Giri",
			age:  40,
		},
		{
			name: "Gri",
			age:  4,
		},
		{
			name: "Gir",
			age:  0,
		},
	}
	for l, v := range u {
		fmt.Println(l, v)
	}
	//fmt.Println(u)
}
