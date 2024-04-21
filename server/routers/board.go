package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shamilsdq/ideaboard-go-svelte/server/dtos"
	"github.com/shamilsdq/ideaboard-go-svelte/server/utils"
)

func GetBoardRouter() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", create).Methods(http.MethodPost)
	return mux
}

func create(w http.ResponseWriter, r *http.Request) {
	var requestDto dtos.CreateBoardRequestDto
	utils.ParseRequestBody(r, &requestDto)

	fmt.Println(requestDto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
