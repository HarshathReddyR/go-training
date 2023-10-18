package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[int]int)
var wg sync.WaitGroup

var m = &sync.Mutex{}

func main() {
	for i := 0; i < 2; i++ {

		wg.Add(1)

		go func(n int) {
			m.Lock()
			defer wg.Done()
			myMap[n] = n * n
			m.Unlock()
		}(i)

	}
	wg.Wait()
	fmt.Println(myMap)
}

// var counter = 0
// var m = &sync.Mutex{}

// func main() {

// 	wg := new(sync.WaitGroup)

// 	for i := 0; i < 3; i++ {

// 		wg.Add(1)

// 		go func() {
// 			m.Lock()
// 			defer wg.Done()

// 			j := counter

// 			time.Sleep(time.Millisecond)

// 			j = j + 1

// 			counter = j
// 			m.Unlock()

// 		}()

// 	}

// 	wg.Wait()

// 	log.Println("counter:", counter)

// }
