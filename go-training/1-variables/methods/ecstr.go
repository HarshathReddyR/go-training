package main

import "fmt"

type User struct {
	Name     string
	Password string
}

var userstroe = map[int]User{} //or var userstore = make(map[int]User)

func (u User) adder(id int) {
	userstroe[id] = u
}
func main() {
	u := User{
		Name:     "abhi",
		Password: "123",
	}
	u1 := User{
		Name:     "Giri",
		Password: "0987",
	}
	u.adder(1)
	u1.adder(2)
	fmt.Println(userstroe)
}
