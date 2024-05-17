package utils

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/shamilsdq/ideaboard-go-svelte/server/dtos"
	"github.com/shamilsdq/ideaboard-go-svelte/server/entities"
)

func HandleIncomingMessages(conn *websocket.Conn, board *entities.Board) {
	defer conn.Close()
	for {
		var dto dtos.BoardSocketDto
		if jsonErr := conn.ReadJSON(&dto); jsonErr != nil {
			continue
		}

		var err error

		switch dto.Code {
		case "POST_CREATE":
			err = handlePostCreate(dto.Content, board)
		case "POST_UPDATE":
			err = handlePostUpdate(dto.Content, board)
		case "POST_DELETE":
			err = handlePostDelete(dto.Content, board)
		default:
			err = fmt.Errorf("unknown request type: %s", dto.Code)
		}

		if err != nil {
			// send json error message to conn
		}
	}
}

func handlePostCreate(content interface{}, board *entities.Board) error {
	var dto dtos.CreatePostDto
	if parseErr := ParseToDto(content, &dto); parseErr != nil {
		return parseErr
	}
	return board.AddPost(dto.SectionId, dto.Content)
}

func handlePostUpdate(content interface{}, board *entities.Board) error {
	var dto dtos.UpdatePostDto
	if parseErr := ParseToDto(content, &dto); parseErr != nil {
		return parseErr
	}
	return board.UpdatePost(dto.Id, dto.SectionId, dto.Content)
}

func handlePostDelete(content interface{}, board *entities.Board) error {
	var dto dtos.DeletePostDto
	if parseErr := ParseToDto(content, &dto); parseErr != nil {
		return parseErr
	}
	return board.DeletePost(dto.Id)
}
