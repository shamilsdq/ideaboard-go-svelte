package routers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetAssetRouter() http.Handler {
	if os.Getenv("DEBUG") == "true" {
		frontendServerUrl, err := url.Parse(os.Getenv("CLIENT_SERVER_URL"))
		if err != nil {
			log.Fatal("Client server URL error", err)
		}
		return httputil.NewSingleHostReverseProxy(frontendServerUrl)
	} else {
		rootPath := os.Getenv("CLIENT_FS_ROOT")
		return &assetRouter{rootPath: rootPath}
	}
}

type assetRouter struct {
	rootPath string
}

func (router *assetRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedFilePath := path.Clean(r.URL.Path)
	fullFilePath := filepath.Join(router.rootPath, path.Clean(r.URL.Path))

	info, err := os.Stat(fullFilePath)

	// Handle File or folder not found
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Handle serving all files, not folders
	if !(info.IsDir()) {
		http.ServeFile(w, r, fullFilePath)
		return
	}

	// Handle static folders
	if strings.HasPrefix(requestedFilePath, "/static") {
		http.NotFound(w, r)
		return
	}

	// Handler non-static folders
	http.ServeFile(w, r, filepath.Join(router.rootPath, "index.html"))
}
