package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type AssetHandler struct {
	rootPath string
}

func (h *AssetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedFilePath := path.Clean(r.URL.Path)
	fullFilePath := filepath.Join(h.rootPath, path.Clean(r.URL.Path))

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
	http.ServeFile(w, r, filepath.Join(h.rootPath, "index.html"))
}
