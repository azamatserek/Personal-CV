package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Experience struct {
	Position     string `json:"position"`
	Organization string `json:"organization"`
	Start        string `json:"start"`
	End          string `json:"end"`
	Description  string `json:"description"`
}

func main() {
	experiences := []Experience{
		{
			Position:     "Associate Professor",
			Organization: "Astana IT University",
			Start:        "Dec 2025",
			End:          "Present",
			Description:  "Teaching, research, leading grants funded by Ministry.",
		},
	}

	http.HandleFunc("/api/experience", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(experiences)
	})

	log.Println("Server running on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
