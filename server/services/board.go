package services

import (
	"sync"

	"github.com/shamilsdq/ideaboard-go-svelte/server/entities"
)

type boardService struct {
	mu     sync.RWMutex
	boards map[string]*entities.Board
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
