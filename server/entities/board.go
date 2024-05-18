package entities

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/shamilsdq/ideaboard-go-svelte/server/dtos"
)

type Board struct {
	title             string
	sections          map[int]*Section
	members           []*websocket.Conn
	posts             map[int]*Post
	mu                sync.RWMutex
	_postIdCounter    int
	_sectionIdCounter int
}

func (board *Board) AddMember(conn *websocket.Conn) {
	board.mu.Lock()
	defer board.mu.Unlock()

	board.members = append(board.members, conn)
}

func (board *Board) RemoveMember(conn *websocket.Conn) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	for i, m := range board.members {
		if m == conn {
			board.members = append(board.members[:i], board.members[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("member not found: %p", conn)
}

func (board *Board) AddPost(sectionId int, content string) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	section, ok := board.sections[sectionId]
	if !ok {
		return fmt.Errorf("section not found: %d", sectionId)
	}

	board._postIdCounter += 1
	newPost := &Post{Content: content, SectionId: sectionId}

	section.AddPostId(board._postIdCounter)
	board.posts[board._postIdCounter] = newPost

	board.broadcast("POST_CREATED", &dtos.PostCreateBroadcastDto{
		Id:        board._postIdCounter,
		SectionId: newPost.SectionId,
		Content:   newPost.Content,
	})

	return nil
}

func (board *Board) UpdatePost(postId int, sectionId int, content string) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	post, postOk := board.posts[postId]
	if !postOk {
		return fmt.Errorf("post not found: %d", postId)
	}

	oldSection := board.sections[post.SectionId]
	oldSection.RemovePostId(postId)

	newSection, sectionOk := board.sections[sectionId]
	if !sectionOk {
		return fmt.Errorf("section not found: %d", sectionId)
	}
	newSection.AddPostId(postId)

	post.Content = content
	post.SectionId = sectionId

	board.broadcast("POST_UPDATED", &dtos.PostUpdateBroadcastDto{
		Id:        postId,
		SectionId: post.SectionId,
		Content:   post.Content,
	})

	return nil
}

func (board *Board) DeletePost(postId int) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	post, postOk := board.posts[postId]
	if !postOk {
		return fmt.Errorf("post not found: %d", postId)
	}

	section, sectionOk := board.sections[post.SectionId]
	if sectionOk {
		section.RemovePostId(postId)
	}

	delete(board.posts, postId)

	board.broadcast("POST_DELETED", &dtos.PostDeleteBroadcastDto{
		Id: postId,
	})

	return nil
}

func (board *Board) AddSection(sectionTitle string) {
	board.mu.Lock()
	defer board.mu.Unlock()

	board._sectionIdCounter += 1
	board.sections[board._sectionIdCounter] = &Section{Title: sectionTitle, PostIds: make([]int, 0)}
}

func (board *Board) UpdateSection(sectionId int, sectionTitle string) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	section, sectionOk := board.sections[sectionId]
	if !sectionOk {
		return fmt.Errorf("section not found: %d", sectionId)
	}

	section.Title = sectionTitle
	return nil
}

func (board *Board) broadcast(code string, content any) {
	dto := &dtos.SocketDto{Code: code, Content: content}
	for _, member := range board.members {
		member.WriteJSON(dto)
	}
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
		title:             title,
		sections:          sectionMap,
		members:           make([]*websocket.Conn, 0),
		posts:             make(map[int]*Post),
		mu:                sync.RWMutex{},
		_postIdCounter:    0,
		_sectionIdCounter: len(sections),
	}
}
