package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"workshop/src/database"
	"workshop/src/types"

	"github.com/julienschmidt/httprouter"
)

func SaveBenefit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user = types.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userId, err := database.DB.InsertUser(user); err != nil {
		log.Printf("Err %s", err.Error())
	} else {
		json.NewEncoder(w).Encode(userId)
	}
}
