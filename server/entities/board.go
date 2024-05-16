package entities

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Board struct {
	title         string
	sections      map[int]*Section
	members       []*websocket.Conn
	posts         map[int]*Post
	mu            sync.RWMutex
	postIdCounter int
}

func NewBoard(title string, sections []string) *Board {
	sectionMap := make(map[int]*Section)
	for idx, sectionTitle := range sections {
		sectionMap[idx] = &Section{
			Title:   sectionTitle,
			PostIds: make([]int, 0),
		}
	}

	return &Board{
		title:         title,
		sections:      sectionMap,
		members:       make([]*websocket.Conn, 0),
		posts:         make(map[int]*Post),
		mu:            sync.RWMutex{},
		postIdCounter: 0,
	}
}
