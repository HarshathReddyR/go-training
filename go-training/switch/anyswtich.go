package main

import "fmt"

func main() {
	checkType("20")
	checkType(User{})
}

type User struct{}

func checkType(i any) {
	switch i.(type) {
	case int:
		fmt.Println("it is an int value")
	case User:
		fmt.Println("it is an user")
	}
}
