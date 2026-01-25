package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./academic.db")
    if err != nil {
       log.Fatal(err)
    }

    // Add the publications table here
    const createTables = `
    CREATE TABLE IF NOT EXISTS experiences (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       position TEXT,
       organization TEXT,
       start TEXT,
       end TEXT,
       description TEXT
    );
    CREATE TABLE IF NOT EXISTS publications (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       title TEXT NOT NULL,
       authors TEXT NOT NULL,
       venue TEXT NOT NULL,
       year INTEGER NOT NULL,
       link TEXT,
       type TEXT -- e.g., "Journal", "Conference"
    );`

    _, err = db.Exec(createTables)
    if err != nil {
       log.Fatal("Failed to create tables:", err)
    }

    return db
}

// Ensure SeedExperiences is still here since main.go calls it
func SeedExperiences(db *sql.DB) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM experiences").Scan(&count)
    if err != nil {
       log.Fatal(err)
    }

    if count == 0 {
       initial := []struct {
          Position, Organization, Start, End, Description string
       }{
          {"Associate Professor", "Astana IT University", "Dec 2025", "Present", "Teaching, research, leading grants funded by Ministry."},
          {"Assistant Professor", "Kazakh-British Technical University", "Sep 2024", "Dec 2025", "Lecturing, curriculum development, student supervision."},
          {"Senior Lecturer", "SDU University", "Sep 2018", "Sep 2024", "Teaching software courses, research projects, mentoring students."},
       }

       for _, exp := range initial {
          _, err := db.Exec(`INSERT INTO experiences(position, organization, start, end, description) VALUES(?,?,?,?,?)`,
             exp.Position, exp.Organization, exp.Start, exp.End, exp.Description)
          if err != nil {
             log.Fatal(err)
          }
       }
    }
}

func SeedPublications(db *sql.DB) {
    var count int
    // Fixed the typo "0 {A" from your snippet
    err := db.QueryRow("SELECT COUNT(*) FROM publications").Scan(&count)
    if err != nil {
        log.Fatal(err)
    }

    if count == 0 {
        _, err := db.Exec(`INSERT INTO publications (title, authors, venue, year, link, type) VALUES 
        ('Deep Learning in IoT', 'John Doe, Jane Smith', 'IEEE Internet of Things Journal', 2025, 'https://doi.org/...', 'Journal'),
        ('Cloud Security Frameworks', 'John Doe', 'International Conference on Software Engineering', 2024, '', 'Conference')`)
        if err != nil {
            log.Fatal("Failed to seed publications:", err)
        }
    }
}