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

	createTableSQL := `CREATE TABLE IF NOT EXISTS experiences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		position TEXT,
		organization TEXT,
		start TEXT,
		end TEXT,
		description TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	return db
}

func SeedExperiences(db *sql.DB) {
	// Check if table is empty
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM experiences").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		// Insert initial data
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
