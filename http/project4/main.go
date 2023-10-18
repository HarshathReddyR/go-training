package main

import (
	"net/http"
	"project4/handers"
)

func main() {

	http.HandleFunc("/home1", handers.Home)
	//start your server
	http.ListenAndServe(":8080", nil)

}
