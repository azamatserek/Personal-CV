package main

import (
    "encoding/json"
    "log"
    "net/http"

    "academic-cv-backend/db"
    "academic-cv-backend/internal/experience" // Ensure this path is correct
)

func main() {
    // 1. Initialize DB (Matches your db.go signature)
    database := db.InitDB()
    defer database.Close()

    // 2. Seed data (Matches your db.go signature)
    db.SeedExperiences(database)

    // 3. Initialize the Service you built
    expService := experience.NewService(database)

    http.HandleFunc("/api/experience", func(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       w.Header().Set("Access-Control-Allow-Origin", "*")

       // Use the service instead of manual SQL here
       experiences := expService.GetAll()

       if err := json.NewEncoder(w).Encode(experiences); err != nil {
          log.Printf("Failed to encode response: %v", err)
       }
    })

    log.Println("Server running on :8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
       log.Fatal(err)
    }
}