package dtos

type BoardCreateRequestDto struct {
	Title    string   `json:"title"`
	Sections []string `json:"sections"`
}

type BoardCreateResponseDto struct {
	BoardId string `json:"boardId"`
}
