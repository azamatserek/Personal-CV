package experience

import (
    "database/sql"
    "log"
)

// ExperienceService now holds a reference to the database connection
type ExperienceService struct {
    db *sql.DB
}

// NewService accepts the db connection from your main/db package
func NewService(db *sql.DB) *ExperienceService {
    return &ExperienceService{
        db: db,
    }
}

// GetAll fetches all rows from the database
func (s *ExperienceService) GetAll() []Experience {
    rows, err := s.db.Query("SELECT position, organization, start, end, description FROM experiences")
    if err != nil {
        log.Printf("Error querying experiences: %v", err)
        return nil
    }
    defer rows.Close()

    var experiences []Experience
    for rows.Next() {
        var exp Experience
        err := rows.Scan(&exp.Position, &exp.Organization, &exp.Start, &exp.End, &exp.Description)
        if err != nil {
            log.Printf("Error scanning experience: %v", err)
            continue
        }
        experiences = append(experiences, exp)
    }

    return experiences
}