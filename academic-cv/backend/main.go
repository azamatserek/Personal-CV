package main

import (
    "encoding/json"
    "net/http"
)

type Publication struct {
    Title string `json:"title"`
    Year  int    `json:"year"`
    DOI   string `json:"doi"`
    Type  string `json:"type"` // Journal / Conference
    Q     string `json:"q"`    // Q1-Q4
}

func publicationsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*") // allow React frontend
    w.Header().Set("Content-Type", "application/json")

    pubs := []Publication{
        {Title: "Crowd Density Estimation...", Year: 2025, DOI: "10.1109/ACCESS.2025.3597393", Type: "Journal", Q: "Q1"},
        {Title: "Few-shot brain tumor classification...", Year: 2025, DOI: "10.11591/eei.v14i5.10706", Type: "Journal", Q: "Q1"},
    }

    json.NewEncoder(w).Encode(pubs)
}


func main() {
    http.HandleFunc("/api/publications", publicationsHandler)
    http.ListenAndServe(":8081", nil) // backend runs on 8081
}
