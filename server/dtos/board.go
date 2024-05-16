package dtos

type CreateBoardRequestDto struct {
	Title    string   `json:"title"`
	Sections []string `json:"sections"`
}

type CreateBoardResponseDto struct {
	BoardId string `json:"boardId"`
}
