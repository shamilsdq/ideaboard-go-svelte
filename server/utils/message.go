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
		if err := conn.ReadJSON(&dto); err != nil {
			fmt.Println("Error reading message: ", err)
			return
		}

		switch dto.Code {
		case "POST_CREATE":
			fmt.Println("Post create request")
		case "POST_UPDATE":
			fmt.Println("Post update request")
		case "POST_DELETE":
			fmt.Println("Post delete request")
		}
	}
}
