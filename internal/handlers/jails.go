package handlers

import (
	"fmt"
	"net/http"

	"github.com/otterlabs/jaildeck/internal/services"
)

type JailHandler struct {
	jailService *services.JailService
}

func NewJailHandler(jailService *services.JailService) *JailHandler {
	return &JailHandler{
		jailService: jailService,
	}
}

func (h *JailHandler) List(w http.ResponseWriter, r *http.Request) {
	jails, err := h.jailService.List(r.Context())
	if err != nil {
		http.Error(w, "failed to list jails", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, "<h1>Jails</h1>")
	fmt.Fprintln(w, "<table>")
	fmt.Fprintln(w, "<tr><th>Name</th></tr>")

	for _, jail := range jails {
		fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", jail.Name, jail.Status)
	}
	fmt.Fprintln(w, "</table>")
}
