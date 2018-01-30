package router

import (
<<<<<<< HEAD
	"io"
=======
>>>>>>> ba2ab54... Changes after PR Review
	"net/http"

	"github.com/jackc/pgx"
)

func isAuthenticated(dbPool *pgx.ConnPool, r *http.Request) bool {
	return true
}

func AuthenticationHandler(dbPool *pgx.ConnPool, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(dbPool, r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}
<<<<<<< HEAD

func LoggingHandler(out io.Writer, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
=======
>>>>>>> ba2ab54... Changes after PR Review