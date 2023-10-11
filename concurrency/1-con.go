package main

import (
	"fmt"
)

func main() {
	go hello() // spawning a goroutine
	fmt.Println("end of the main")
	//time.Sleep(2 * time.Second)
}

func hello() {
	fmt.Println("hello from the hello func")
}
