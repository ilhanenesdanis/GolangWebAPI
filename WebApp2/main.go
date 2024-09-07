package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstName"
	LastName  string "json:lastName"
	Age       int    "json:age"
}

func main() {

	apiRoot := "/api"

	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API HOME"}
		output, err := json.Marshal(message)
		checkError(err)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "İlhan Enes", LastName: "Daniş", Age: 23},
			User{ID: 2, FirstName: "Rümeysa", LastName: "Daniş", Age: 17},
			User{ID: 3, FirstName: "Ramazan", LastName: "Daniş", Age: 52},
		}

		message := users

		output, err := json.Marshal(message)

		checkError(err)

		fmt.Fprint(w, string(output))

	})

	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 1, FirstName: "İlhan Enes", LastName: "Daniş", Age: 23}
		message := user

		output, err := json.Marshal(message)

		checkError(err)

		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":9000", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}
