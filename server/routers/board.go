package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shamilsdq/ideaboard-go-svelte/server/dtos"
	"github.com/shamilsdq/ideaboard-go-svelte/server/services"
	"github.com/shamilsdq/ideaboard-go-svelte/server/utils"
)

func GetBoardRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/boards", create).Methods(http.MethodPost)
	return router
}

func create(w http.ResponseWriter, r *http.Request) {
	var requestDto dtos.CreateBoardRequestDto
	utils.ParseRequestBody(r, &requestDto)

	boardId, createErr := services.
		GetBoardService().
		CreateBoard(requestDto.Title, requestDto.Sections)
	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseDto := dtos.CreateBoardResponseDto{BoardId: boardId}
	responseJson, jsonErr := json.Marshal(responseDto)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}
