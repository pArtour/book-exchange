package handler

import (
	"net/http"
)

func (h *Handler) HandleBooksGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(ctxKeyUser).(int)
		h.respond(w, http.StatusOK, map[string]int{"id": id})
	}
}
