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
	router.PathPrefix("/").Handler(routers.GetAssetRouter()).Methods(http.MethodGet)
	return handlers.LoggingHandler(os.Stdout, router)
}
