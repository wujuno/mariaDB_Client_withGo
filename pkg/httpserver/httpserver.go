package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Learning REST API with mariaDB and go")
	}).Methods(http.MethodGet)

	/* r.HandleFunc("idol-list", idolListHandler).Methods(http.MethodGet) */

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080",handler)
}