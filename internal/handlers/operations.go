package handlers

import (
	"log"
	"net/http"

	"github.com/ottermq/jaildeck/internal/operations"
	"github.com/ottermq/jaildeck/internal/services"
	"github.com/ottermq/jaildeck/internal/views"
)

type OperationHandler struct {
	service  *services.OperationService
	renderer *views.Renderer
}

func NewOperationHandler(service *services.OperationService, renderer *views.Renderer) *OperationHandler {
	return &OperationHandler{service: service, renderer: renderer}
}

func (h *OperationHandler) List(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.Recent(r.Context(), 50)
	if err != nil {
		http.Error(w, "failed to list operations", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title  string
		Entries []operations.Entry
	}{
		Title:  "Operations",
		Entries: entries,
	}

	if err := h.renderer.Render(w, "operations", data); err != nil {
		log.Printf("failed to render page: %s", err.Error())
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}
