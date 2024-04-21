package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewServer() http.Handler {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(GetAssetServer()).Methods(http.MethodGet)
	return handlers.LoggingHandler(os.Stdout, router)
}

func GetAssetServer() http.Handler {
	if os.Getenv("DEBUG") == "true" {
		frontendServerUrl := &url.URL{Host: os.Getenv("CLIENT_SERVER_URL")}
		return httputil.NewSingleHostReverseProxy(frontendServerUrl)
	} else {
		rootPath := os.Getenv("CLIENT_FS_ROOT")
		return &AssetHandler{rootPath: rootPath}
	}
}
