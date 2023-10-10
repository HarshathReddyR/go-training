// package main
// type student struct {
// }

// func (s student) Write(p []byte) (n int, err error) {
//     //TODO implement me
//     panic("implement me")
// }

// func (s student) Read(p []byte) (n int, err error) {
//     //TODO implement me
//     panic("implement me")
// }

//	func main() {
//	    var s student
//	    var r io.Reader = s
//	    var w io.Writer = s
//	    var rw io.ReadWriter = s
//	    _, _, _ = r, w, rw
//	}
package main

import "fmt"

type User struct {
	Name string
}
type Walk interface {
	walk(string)
}
type Run interface {
	run(string)
}
type WalkRun interface {
	Walk
	Run
}

func (u User) walk(s string) {
	fmt.Println(u.Name, s)
}
func (u User) run(s string) {
	fmt.Println(u.Name, s)
}

// func (u User)
func main() {
	u := User{
		Name: "Abhi",
	}
	// u1 := User{
	// 	Name: "Abhi",
	// }
	// u.run("is for Fitness")
	var wr WalkRun = u
	//w.walk("is running")
	wr.run("is running")
}
