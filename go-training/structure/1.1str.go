package main

import "fmt"

//creating a structure
type user struct {
	Name  string
	Age   int
	Marks []int
}

func main() {
	//using structure
	u := user{
		Name:  "person 1",
		Age:   20,
		Marks: []int{10, 20, 30, 40, 50},
	}
	fmt.Println(u.Name) //getting the structure values
	fmt.Println(u.Age)
	fmt.Println(u.Marks)
	fmt.Printf("%+v", u)
	fmt.Printf("%#v", u)
}
