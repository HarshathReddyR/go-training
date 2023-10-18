// package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"project/models"
// 	"strconv"
// )

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	//set content type
// 	w.Header().Set("context-text", "application/json")
// 	userIdString := r.URL.Query().Get("user_id")
// 	//convert this to uint (strconv.ParseUint)
// 	id, err := strconv.ParseUint(userIdString, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		conErr := map[string]string{"Message": "conversion error ", "err": err.Error()}
// 		j, err := json.Marshal(conErr)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
// 			return
// 		}
// 		w.Write(j)

// 		return
// 	}
// 	userdata, err := models.FetchUser(id)
// 	if err != nil {
// 		log.Println(err)
// 		fecErr := map[string]string{"Message": "user not found ", "err": err.Error()}
// 		fj, err := json.Marshal(fecErr)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
// 			return
// 		}
// 		w.Write(fj)
// 		return
// 	}
// 	userjason, err := json.Marshal(userdata)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	w.Write(userjason)

// 	// if error happens report to the client in json // write status code

// 	//appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

// 	//call the FetchUser function from the models package

// 	// write to the client with user data or error

// }

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/models"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	//set content type

	userIdString := r.URL.Query().Get("user_id") // this will fetch the

	//convert this to uint (strconv.ParseUint)
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	// if error happens report to the client in json // write status code
	if err != nil {

		log.Println("Error: ", err)

		errorInConversion := map[string]string{"msg": "not a valid number"}

		jsonData, err := json.Marshal(errorInConversion)

		if err != nil {
			log.Println("Error while converting error to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusBadRequest)

		w.Write(jsonData)

		return

	}

	//appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

	//call the FetchUser function from the models package

	uData, err := models.FetchUser(userId)

	if err != nil {

		fetchError := map[string]string{"msg": "user not found"}

		errData, err := json.Marshal(fetchError)

		if err != nil {

			log.Println("Error while parsing fetchuser error conversion: ", err)

			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

			return

		}

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(errData)

		return

	}

	// write to the client with user data or error

	userData, err := json.Marshal(uData)

	if err != nil {

		log.Println("Error while converting user data to json", err)

		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

		return

	}

	w.Write(userData)

}
