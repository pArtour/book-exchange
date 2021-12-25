package handler

import (
	"encoding/json"
	"errors"
	"github.com/pArtour/book-exchange/pkg/model"
	"net/http"
)

var (
	errInvalidEmailOrPassword = errors.New("Incorrect password or email")
	errNotAuthenticated       = errors.New("Not authenticated")
)

func (h *Handler) HandleUserCreate() http.HandlerFunc {
	request := &model.User{}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			h.error(w, http.StatusBadRequest, err)
			return
		}

		if err := h.services.Auth.CreateUser(request); err != nil {
			h.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		h.respond(w, http.StatusCreated, map[string]string{"status": "created"})
	}
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleUserSignIn() http.HandlerFunc {
	request := &user{}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			h.error(w, http.StatusBadRequest, err)
			return
		}

		token, err := h.services.Auth.GenerateToken(request.Email, request.Password)
		if err != nil {
			h.error(w, http.StatusBadRequest, errInvalidEmailOrPassword)
			return
		}

		h.respond(w, http.StatusCreated, map[string]string{"token": token})
	}
}
