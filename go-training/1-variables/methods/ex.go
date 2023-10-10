package main

import "fmt"

type Lang struct {
	LangName string
}
//for storing
var store map[int]Lang{}

func (l Lang)adder(s string){
	store[s]=
}
func (l *Lang) update(lang string) {
	l.LangName = lang
}
func main() {
	l := Lang{
		LangName: "kannada",
	}
	fmt.Println(l) //before updations
	l.update("English")
	fmt.Println(l) //after updated
}
