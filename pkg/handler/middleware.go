package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	authHeader        = "Authorization"
	ctxKeyUser ctxKey = iota
)

type ctxKey int8

func (h *Handler) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			h.error(w, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			h.error(w, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		userId, err := h.services.Auth.ParseToken(headerParts[1])
		if err != nil {
			h.error(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, userId)))
	})
}

//const (
//	ctxKeyRequestID
//)
func (h *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := h.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			//"request_id":  r.Context().Value(ctxKeyRequestID),
		})

		logger.Infof("started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		logger.Infof(
			"completed with %d %s in %v ",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}
