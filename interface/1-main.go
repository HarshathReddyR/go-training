package main

import "fmt"

type information interface {
	Info()
}
type Zamp struct {
	path string
}

type Logrus struct {
	path string
}

func Dolog(i information) {
	i.Info()
}
func (z Zamp) Info() {
	fmt.Println(z.path)
}
func (l Logrus) Info() {
	fmt.Println(l.path)
}
func main() {
	z := Zamp{
		path: "Terminal",
	}
	l := Logrus{
		path: "file",
	}
	Dolog(z)
	Dolog(l)
}
