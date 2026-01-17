package experience

import (
    "encoding/json"
    "net/http"
)

type Handler struct {
    service *ExperienceService
}

func NewHandler(s *ExperienceService) *Handler {
    return &Handler{service: s}
}

func (h *Handler) GetExperiences(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") // allow React frontend

    json.NewEncoder(w).Encode(h.service.GetAll())
}
