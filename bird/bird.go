package bird

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type birdStruct struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var Birds []birdStruct

func GetBirdHandler(w http.ResponseWriter, r *http.Request) {
	birdListBytes, err := json.Marshal(Birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := birdStruct{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	Birds = append(Birds, bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}