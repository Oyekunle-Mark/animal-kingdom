package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/oyekunle-mark/animal-kingdom/bird"
	"github.com/oyekunle-mark/animal-kingdom/store"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	fmt.Println("Starting server...")

	connString := "dbname=template1 sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	store.InitStore(store.DbStore{db: db})

	r := newRouter()

	fmt.Println("Serving on port 3000")
	http.ListenAndServe(":3000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
