package dtos

type BoardCreateRequestDto struct {
	Title    string   `json:"title" validate:"required,min=1"`
	Sections []string `json:"sections" validate:"required,gt=0"`
}

type BoardCreateResponseDto struct {
	BoardId string `json:"boardId"`
}

type BoardDataDto struct {
	Title       string `json:"title"`
	Sections    []*SectionDataDto
	Posts       []*PostDataDto
	MemberCount int
}
