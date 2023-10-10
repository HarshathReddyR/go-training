package main

import (
	"fmt"
	//"io"
	"log"
	//"github.com/sirupsen/logrus/hooks/writer"
)

type user struct {
	name string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Println("writing in the user's")
	return n, err
}

func main() {
	u := user{name: "ajay"}
	l := log.New(u, "sales:", log.Lshortfile)
	l.Println()
	//io.Writer()

}
