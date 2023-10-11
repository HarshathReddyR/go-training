package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "1"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "2"
	}()
	go func() {
		c3 <- "3"
	}()
	fmt.Println(<-c1)
	fmt.Println(<-c3)
	fmt.Println(<-c1)
}
