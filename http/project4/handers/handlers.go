package handers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	UsreId  string
	Hobbies []string
}

//	func Home(rw http.ResponseWriter, rq *http.Request) {
//		rw.WriteHeader(http.StatusBadRequest)
//		rw.Write([]byte("Hi hello this by using packages"))
//	}
func Home(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-text", "application/json")

	u := User{
		UsreId: "Abhi",
		Hobbies: []string{
			"Playing",
			"Sleeping",
		},
	}
	j, err := json.Marshal(u)
	rw.WriteHeader(http.StatusOK)
	if err != nil {
		fmt.Println(err)
		return
	}
	rw.Write(j)
}
