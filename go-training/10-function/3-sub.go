package main

import "fmt"

func get(s string) {
	fmt.Println("get" + s)
}
func post(s string) {
	fmt.Println("post" + s)
}
func put(s string) {
	fmt.Println("put" + s)
}
func request(r func(s string), endpoint string) {
	r(endpoint)
}
func main() {
	request(get, "hh")
}
