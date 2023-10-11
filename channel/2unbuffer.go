package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	var r []int
	ch := make(chan int) //unbuffered channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 10
		ch <- 20 // send will block until no receiver is ready
	}()

	r := <-ch, <-ch //it is a blocking call until we don't recv the value
	// x1 := <-ch
	fmt.Println(x)
	fmt.Println(x1)
	wg.Wait()
	fmt.Println("end of main")

}
