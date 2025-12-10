package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"Feast-Serve/menu"
)

func generateMenuHandler(w http.ResponseWriter, r *http.Request) {
	items, err := menu.LoadMenuFromJSON("./data/master_menu.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plan := menu.GenerateMenuSuggestions(items, 7, 3, 550, 800)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/generate-menu", generateMenuHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Used only when running locally
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
