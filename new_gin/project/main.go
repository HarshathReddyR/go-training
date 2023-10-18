package main

import (
	"net/http"
	"project/handlers"
)

func main() {
	http.HandleFunc("/user", handlers.GetUser)
	//start your server
	http.ListenAndServe(":8085", nil)
}
