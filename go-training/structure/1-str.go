package main

import (
	"fmt"
)

type money int //uinderlaying type("we are creating our own type")

func main() {
	var Rupee money = 10
	fmt.Printf("%T", Rupee)
	//i := 120
	//Rupee = i
	//fmt.Println(Rupee)
}
