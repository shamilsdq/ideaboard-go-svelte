package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewServer() http.Handler {
	router := mux.NewRouter()
	return handlers.LoggingHandler(os.Stdout, router)
}
