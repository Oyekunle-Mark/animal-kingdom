package main

import (
	"fmt"
	"net/http"
	"github.com/oyekunle-mark/animal-kingdom/bird"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", bird.GetBirdHandler).Methods("GET")
	r.HandleFunc("/bird", bird.CreateBirdHandler).Methods("POST")
	
	return r
}

func main() {
	r := newRouter()

	http.ListenAndServe(":3000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
