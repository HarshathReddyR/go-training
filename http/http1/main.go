package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/home", home)
	//start your server
	http.ListenAndServe(":8080", nil)

}

func home(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("Hi hello"))
}
