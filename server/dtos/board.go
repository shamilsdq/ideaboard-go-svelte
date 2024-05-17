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

type CreatePostDto struct {
	SectionId int    `json:"sectionId"`
	Content   string `json:"content"`
}

type UpdatePostDto struct {
	Id        int    `json:"id"`
	SectionId int    `json:"sectionId"`
	Content   string `json:"content"`
}

type DeletePostDto struct {
	Id int `json:"id"`
}
