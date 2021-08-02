package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"alive": true})
}

func getPosts(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homepage)
	r.HandleFunc("/getposts", getPosts)

	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), handlers.AllowedOrigins([]string{"localhost:3000"}))(r)))
}
