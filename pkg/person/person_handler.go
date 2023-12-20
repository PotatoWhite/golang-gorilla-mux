package person

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const PersonPath = "/persons/{id}"

func PersonHandler(w http.ResponseWriter, r *http.Request) {

	// id from path
	pathVars := mux.Vars(r)
	id, err := strconv.Atoi(pathVars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, _ := json.Marshal(personsTable[id])
	w.Write(jsonResponse)
}
