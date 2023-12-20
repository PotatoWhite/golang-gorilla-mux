package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest-api-basic/pkg/person"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc(person.PersonsPath, person.PersonsHandler).Methods("GET")
	router.HandleFunc(person.PersonPath, person.PersonHandler).Methods("GET")
	router.HandleFunc(person.CreatePersonPath, person.CreatePersonHandler).Methods("POST")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
