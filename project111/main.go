// package main

// import "fmt"

// func Add(s int) int{
// 	fmt.Println("in Add")
// 	return Sub(1,2)

// }
// func Sub(x int, y int)int {
// 	fmt.Println("in Sub")
// 	return x-y
// }
// func Op(a int) int{
// 	fmt.Println("in Op")
// 	return Add(a)
// }
// func main() {
// 	Op

// }
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", Mid(Mid1(HomePage)))
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page invoked")
	fmt.Fprintln(w, "this is my home")
}
func Mid1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware invoked")
		next(w, r)
	}
}
func Mid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware invoked")
		next(w, r)
	}
}
