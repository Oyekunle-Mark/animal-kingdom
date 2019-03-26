package bird

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/oyekunle-mark/animal-kingdom/store"
)

var sqlStore store.Store

func GetBirdHandler(w http.ResponseWriter, r *http.Request) {
	birds, err := sqlStore.GetBirds()

	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := store.Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	err = sqlStore.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}