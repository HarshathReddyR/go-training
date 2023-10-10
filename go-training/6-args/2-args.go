package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println(" somthing from main")
}

func greet() {
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("please provide name , age , marks")
		os.Exit(1)
		//return // stops the exec of the current func
	}
	name := data[0]
	ageString := data[1]

	age, err := strconv.Atoi(ageString)
	if err != nil {
		// there is some kind of error
		log.Println(err)
		os.Exit(1)
		//return //its stops the current function
	}

	//marksString := data[2]
	marksString := data[2]
	marks, err := strconv.Atoi(marksString) //here we update the err which is used before
	if err != nil {
		// there is some kind of error
		log.Println(err)
		//return
		os.Exit(1)
	}
	fmt.Println(name, age, marks)
}
