package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pArtour/book-exchange/pkg/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	router   *mux.Router
	services *service.Service
	logger   *logrus.Logger
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{
		router:   mux.NewRouter(),
		services: services,
		logger:   logrus.New(),
	}

	h.ConfigRouter()
	return h
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.router.ServeHTTP(writer, request)
}

func (h *Handler) ConfigRouter() {
	h.router.HandleFunc("/sign-up", h.HandlerUserCreate()).Methods("GET")
	//h.router.HandleFunc("/sign-in").Methods("POST")
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	h.respond(w, statusCode, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
