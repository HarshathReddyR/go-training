package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer recoveryFunc()
	doSomething()
	fmt.Println("end of the main, stopped gracefully")
}

func doSomething() {

	var i []int
	i[100] = 1000
}

func recoveryFunc() {
	msg := recover()
	if msg != nil {
		fmt.Println(string(debug.Stack()))
		fmt.Println(msg)
		fmt.Println("recovered from the panic")
	}
}
