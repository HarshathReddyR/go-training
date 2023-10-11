package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	//wg.Add(10)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg)
	}
	wg.Wait()
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	go func(id int) {
		defer wg.Done()
		fmt.Println("work", id)
	}(i)

}
