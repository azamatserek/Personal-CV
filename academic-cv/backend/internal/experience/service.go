package experience

// ExperienceService provides business logic
type ExperienceService struct {
    experiences []Experience
}

// NewService initializes the service
func NewService() *ExperienceService {
    return &ExperienceService{
        experiences: []Experience{
            {Position: "Associate Professor", Organization: "Astana IT University", Start: "Dec 2025", End: "Present", Description: "Teaching, research, leading grants funded by Ministry."},
            {Position: "Assistant Professor", Organization: "Kazakh-British Technical University", Start: "Sep 2024", End: "Dec 2025", Description: "Lecturing, curriculum development, student supervision."},
            {Position: "Senior Lecturer", Organization: "SDU University", Start: "Sep 2018", End: "Sep 2024", Description: "Teaching software courses, research projects, mentoring students."},
        },
    }
}

// GetAll returns all experiences
func (s *ExperienceService) GetAll() []Experience {
    return s.experiences
}
