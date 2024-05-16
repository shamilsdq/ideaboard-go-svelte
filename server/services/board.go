package services

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/shamilsdq/ideaboard-go-svelte/server/entities"
	"github.com/shamilsdq/ideaboard-go-svelte/server/utils"
)

type boardService struct {
	mu     sync.RWMutex
	boards map[string]*entities.Board
}

func (bs *boardService) CreateBoard(title string, sections []string) (string, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	newBoardId := bs.generateNewBoardId()
	bs.boards[newBoardId] = entities.NewBoard(title, sections)
	return newBoardId, nil
}

func (bs *boardService) DeleteBoard(boardId string) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if _, ok := bs.boards[boardId]; !ok {
		return fmt.Errorf("board does not exist: %s", boardId)
	}

	delete(bs.boards, boardId)
	return nil
}

func (bs *boardService) GetBoard(boardId string) (*entities.Board, error) {
	if board, ok := bs.boards[boardId]; ok {
		return board, nil
	} else {
		return nil, fmt.Errorf("board does not exist: %s", boardId)
	}
}

func (bs *boardService) generateNewBoardId() string {
	for {
		id := fmt.Sprintf("%s-%s-%s",
			utils.GenerateRandomString(3),
			utils.GenerateRandomString(4),
			utils.GenerateRandomString(3))
		if _, ok := bs.boards[id]; !ok {
			return id
		}
	}
}

func (bs *boardService) JoinBoard(boardId string, conn *websocket.Conn) error {
	board, ok := bs.boards[boardId]
	if !ok {
		return fmt.Errorf("board does not exist: %s", boardId)
	}
	board.AddMember(conn)
	conn.SetCloseHandler(func(code int, text string) error {
		board.RemoveMember(conn)
		return nil
	})
	return nil
}

var once sync.Once
var instance *boardService

func GetBoardService() *boardService {
	once.Do(func() {
		instance = &boardService{
			mu:     sync.RWMutex{},
			boards: make(map[string]*entities.Board),
		}
	})
	return instance
}
