package publications

import "database/sql"

type Service struct {
    DB *sql.DB
}

func (s *Service) GetAll() ([]Publication, error) {
    rows, err := s.DB.Query("SELECT id, title, authors, venue, year, link, type FROM publications ORDER BY year DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var pubs []Publication
    for rows.Next() {
        var p Publication
        if err := rows.Scan(&p.ID, &p.Title, &p.Authors, &p.Venue, &p.Year, &p.Link, &p.Type); err != nil {
            return nil, err
        }
        pubs = append(pubs, p)
    }
    return pubs, nil
}