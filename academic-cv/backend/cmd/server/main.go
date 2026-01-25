package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"

    "academic-cv-backend/db"
    "academic-cv-backend/internal/experience"
    "academic-cv-backend/internal/publications"
)

// Simple CORS middleware to handle headers for all routes
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Access-Control-Allow-Origin", "*")
       w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
       w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

       if r.Method == "OPTIONS" {
          w.WriteHeader(http.StatusOK)
          return
       }

       next(w, r)
    }
}

func main() {
    // Initialize DB
    database := db.InitDB()
    defer database.Close()

    // Seed data for both tables
    db.SeedExperiences(database)
    db.SeedPublications(database)

    // Initialize Services
    expService := experience.NewService(database)
    pubService := publications.NewService(database)
    pubHandler := publications.NewHandler(pubService)

    // ========================================
    // PUBLIC API ROUTES
    // ========================================

    // Experience Route
    http.HandleFunc("/api/experience", enableCORS(func(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       experiences := expService.GetAll()
       json.NewEncoder(w).Encode(experiences)
    }))

    // Publications API Route (GET all publications)
    http.HandleFunc("/api/publications", enableCORS(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            pubHandler.GetPublications(w, r)
        case http.MethodPost:
            pubHandler.CreatePublication(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }))

    // Delete Publication Route (DELETE /api/publications/{id})
    http.HandleFunc("/api/publications/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasPrefix(r.URL.Path, "/api/publications/") && r.Method == http.MethodDelete {
            pubHandler.DeletePublication(w, r)
        } else {
            http.Error(w, "Not found", http.StatusNotFound)
        }
    }))

    // ========================================
    // ADMIN ROUTES
    // ========================================

    // Admin page for managing publications
    http.HandleFunc("/admin/publications", pubHandler.ServeAdminPage)

    // ========================================
    // START SERVER
    // ========================================

    log.Println("========================================")
    log.Println("Server running on :8081")
    log.Println("Admin Panel: http://localhost:8081/admin/publications")
    log.Println("API Endpoint: http://localhost:8081/api/publications")
    log.Println("========================================")

    if err := http.ListenAndServe(":8081", nil); err != nil {
       log.Fatal(err)
    }
}