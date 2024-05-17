package dtos

type CreateBoardRequestDto struct {
	Title    string   `json:"title"`
	Sections []string `json:"sections"`
}

type CreateBoardResponseDto struct {
	BoardId string `json:"boardId"`
}

type BoardSocketDto struct {
	Code    string `json:"code"`
	Content any    `json:"content"`
}

type BoardSocketErrorDto struct {
	Error string `json:"error"`
}
