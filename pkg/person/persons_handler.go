package person

import (
	"encoding/json"
	"net/http"
)

const PersonsPath = "/persons"

func PersonsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(personsTable)
	w.Write(jsonResponse)
}
