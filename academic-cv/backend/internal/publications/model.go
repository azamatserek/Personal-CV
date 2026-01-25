package publications

type Publication struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Authors string `json:"authors"`
    Venue   string `json:"venue"`
    Year    int    `json:"year"`
    Link    string `json:"link"`
    Type    string `json:"type"`
}