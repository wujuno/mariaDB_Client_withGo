package httpserver

import (
	"encoding/json"
	"mariadbClient/db"
	"net/http"
)



func idolListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	members, err := db.SelectIdolInfo()
	if err != nil {
		http.Error(w, "Failed to fetch idol information", http.StatusInternalServerError )
		return
	}

	err = json.NewEncoder(w).Encode(members)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}