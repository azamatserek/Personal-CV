package publications

import (
    "encoding/json"
    "net/http"
)

type Handler struct {
    Service *Service
}

func (h *Handler) GetPublications(w http.ResponseWriter, r *http.Request) {
    pubs, err := h.Service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(pubs)
}