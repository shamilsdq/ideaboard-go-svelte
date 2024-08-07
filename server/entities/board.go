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

	board.broadcast("MEMBER_JOINED", &dtos.MemberJoinBroadcastDto{
		MemberCount: len(board.members) + 1,
	})

	board.members = append(board.members, conn)

	boardData := board.generateBoardData()
	conn.WriteJSON(&dtos.SocketDto{Code: "INITIAL", Content: boardData})
}

func (board *Board) RemoveMember(conn *websocket.Conn) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	for i, m := range board.members {
		if m == conn {
			board.members = append(board.members[:i], board.members[i+1:]...)
			board.broadcast("MEMBER_EXITED", &dtos.MemberExitBroadcastDto{
				MemberCount: len(board.members),
			})
			return nil
		}
	}

	return fmt.Errorf("member not found: %p", conn)
}

func (board *Board) AddPost(sectionId int, content string) error {
	board.mu.Lock()
	defer board.mu.Unlock()

	_, ok := board.sections[sectionId]
	if !ok {
		return fmt.Errorf("section not found: %d", sectionId)
	}

	board._postIdCounter += 1
	newPost := &Post{Content: content, SectionId: sectionId}

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

	_, sectionOk := board.sections[sectionId]
	if !sectionOk {
		return fmt.Errorf("section not found: %d", sectionId)
	}

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

	_, postOk := board.posts[postId]
	if !postOk {
		return fmt.Errorf("post not found: %d", postId)
	}

	delete(board.posts, postId)

	board.broadcast("POST_DELETED", &dtos.PostDeleteBroadcastDto{Id: postId})

	return nil
}

func (board *Board) AddSection(sectionTitle string) {
	board.mu.Lock()
	defer board.mu.Unlock()

	board._sectionIdCounter += 1
	board.sections[board._sectionIdCounter] = &Section{Title: sectionTitle}
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

func (board *Board) generateBoardData() *dtos.BoardDataDto {
	sectionDtos := make([]*dtos.SectionDataDto, 0)
	for sectionId, section := range board.sections {
		sectionDtos = append(sectionDtos, &dtos.SectionDataDto{
			Id:    sectionId,
			Title: section.Title,
		})
	}

	postDtos := make([]*dtos.PostDataDto, 0)
	for postId, post := range board.posts {
		postDtos = append(postDtos, &dtos.PostDataDto{
			Id:        postId,
			SectionId: post.SectionId,
			Content:   post.Content,
		})
	}

	return &dtos.BoardDataDto{
		Title:       board.title,
		Sections:    sectionDtos,
		Posts:       postDtos,
		MemberCount: len(board.members),
	}
}

func NewBoard(title string, sections []string) *Board {
	sectionMap := make(map[int]*Section)
	for idx, sectionTitle := range sections {
		sectionMap[idx+1] = &Section{Title: sectionTitle}
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
