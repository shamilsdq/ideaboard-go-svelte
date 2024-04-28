package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shamilsdq/ideaboard-go-svelte/server/routers"
)

func NewServer() http.Handler {
	router := mux.NewRouter()

	router.Use(corsMiddleware)

	router.PathPrefix("/").Handler(routers.GetAssetRouter()).Methods(http.MethodGet)
	router.PathPrefix("/api").Handler(routers.GetBoardRouter())

	return handlers.LoggingHandler(os.Stdout, router)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
