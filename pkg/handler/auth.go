package handler

import "net/http"

func (h *Handler) HandlerUserCreate() http.HandlerFunc {
	type User struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}
}
