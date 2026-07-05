package handlers

import (
	"net/http"

	"github.com/otterlabs/jaildeck/internal/services"
	"github.com/otterlabs/jaildeck/internal/views"
)

type JailHandler struct {
	service  *services.JailService
	renderer *views.Renderer
}

func NewJailHandler(jailService *services.JailService, renderer *views.Renderer) *JailHandler {
	return &JailHandler{
		service:  jailService,
		renderer: renderer,
	}
}

func (h *JailHandler) List(w http.ResponseWriter, r *http.Request) {
	jails, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, "failed to list jails", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Jails any
	}{
		Title: "Jails",
		Jails: jails,
	}

	if err := h.renderer.Render(w, "pages/jails.html", data); err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}
