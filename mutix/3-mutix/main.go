// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	var sm sync.Map

// 	// Store key-value pairs in the sync.Map
// 	sm.Store("color", "green")
// 	v, ok := sm.Load("abc")
// 	if !ok {
// 		fmt.Println("value is not there")
// 		return
// 	}

//		//LoadOrStore returns the existing value for the key if present.
//		//Otherwise, it stores and returns the given value
//		if v, loaded := sm.LoadOrStore("color", "green"); loaded {
//			fmt.Println("Found existing value:", v) // Prints: "Found existing value: orange"
//		} else {
//			fmt.Println("Stored new value:", v) // This won't run in this case
//		}
//		fmt.Println(v)
//	}
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initialize a new map
	var m sync.Map

	// Store: Set key/value pairs
	m.Store("animal", "giraffe")
	m.Store("color", "orange")

	// Load: Get a value
	if v, ok := m.Load("animal"); ok {
		fmt.Println("Loaded value:", v) // Prints: "Loaded value: giraffe"
	}

	// LoadOrStore: Returns existing value for a key or store a new value
	if v, loaded := m.LoadOrStore("color", "green"); loaded {
		fmt.Println("Found existing value:", v) // Prints: "Found existing value: orange"
	} else {
		fmt.Println("Stored new value:", v) // This won't run in this case
	}

	// Delete: Remove a key/value pair
	m.Delete("animal")

	// Use Load to show that "animal" was deleted
	if _, ok := m.Load("animal"); !ok {
		fmt.Println("Value not found") // Prints: "Value not found"
	}

	// Range: Iterating over the map
	m.Range(func(k, v interface{}) bool {
		fmt.Println("Key:", k, "Value:", v)
		return true
	})
}
