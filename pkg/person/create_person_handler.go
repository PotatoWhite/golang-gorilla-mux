package person

import (
	"encoding/json"
	"net/http"
)

const CreatePersonPath = "/persons"

func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {

	// request body
	aPerson := Person{}
	json.NewDecoder(r.Body).Decode(&aPerson)

	// add to table
	aPerson.Id = len(personsTable)
	personsTable = append(personsTable, aPerson)

	// response body
	w.WriteHeader(http.StatusCreated)
	jsonResponse, _ := json.Marshal(aPerson)
	w.Write(jsonResponse)
}
