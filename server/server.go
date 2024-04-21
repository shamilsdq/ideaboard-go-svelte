package server

import "net/http"

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
