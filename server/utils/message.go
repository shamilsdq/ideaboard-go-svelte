package utils

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/shamilsdq/ideaboard-go-svelte/server/dtos"
	"github.com/shamilsdq/ideaboard-go-svelte/server/entities"
)

func HandleIncomingMessages(conn *websocket.Conn, board *entities.Board) {
	defer Cleanup(conn, board)

	for {
		var dto dtos.SocketDto
		if readErr := conn.ReadJSON(&dto); readErr != nil {
			fmt.Printf("JSON read error: %v\n", readErr)
			break
		}

		var errors []string

		switch dto.Code {
		case "POST_CREATE":
			errors = handlePostCreate(dto.Content, board)
		case "POST_UPDATE":
			errors = handlePostUpdate(dto.Content, board)
		case "POST_DELETE":
			errors = handlePostDelete(dto.Content, board)
		default:
			errors = []string{fmt.Sprintf("unknown request type: %s", dto.Code)}
		}

		if errors != nil {
			errDto := dtos.SocketErrorDto{Errors: errors}
			conn.WriteJSON(&errDto)
		}
	}
}

func handlePostCreate(content interface{}, board *entities.Board) []string {
	var dto dtos.PostCreateDto
	if parseErr := ParseAndValidate(content, &dto); parseErr != nil {
		return parseErr
	}
	if err := board.AddPost(dto.SectionId, dto.Content); err != nil {
		return []string{err.Error()}
	}
	return nil
}

func handlePostUpdate(content interface{}, board *entities.Board) []string {
	var dto dtos.PostUpdateDto
	if parseErr := ParseAndValidate(content, &dto); parseErr != nil {
		return parseErr
	}
	if err := board.UpdatePost(dto.Id, dto.SectionId, dto.Content); err != nil {
		return []string{err.Error()}
	}
	return nil
}

func handlePostDelete(content interface{}, board *entities.Board) []string {
	var dto dtos.PostDeleteDto
	if parseErr := ParseAndValidate(content, &dto); parseErr != nil {
		return parseErr
	}
	if err := board.DeletePost(dto.Id); err != nil {
		return []string{err.Error()}
	}
	return nil
}
