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
	h.router.Use(h.logRequest)
	//h.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	h.router.HandleFunc("/sign-up", h.HandleUserCreate()).Methods("POST")
	h.router.HandleFunc("/sign-in", h.HandleUserSignIn()).Methods("POST")
	h.router.Use(h.authenticate)
	h.router.HandleFunc("/books", h.HandleBooksGet()).Methods("GET")
}

func (h *Handler) error(w http.ResponseWriter, statusCode int, err error) {
	h.respond(w, statusCode, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
