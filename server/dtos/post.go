package dtos

type PostCreateDto struct {
	SectionId int    `json:"sectionId" validate:"required,gt=0"`
	Content   string `json:"content" validate:"required,min=1"`
}

type PostUpdateDto struct {
	Id        int    `json:"id" validate:"required,gt=0"`
	SectionId int    `json:"sectionId" validate:"required,gt=0"`
	Content   string `json:"content" validate:"required,min=1"`
}

type PostDeleteDto struct {
	Id int `json:"id" validate:"required,gt=0"`
}

type PostCreateBroadcastDto = PostUpdateDto
type PostUpdateBroadcastDto = PostUpdateDto

type PostDeleteBroadcastDto = PostDeleteDto

type PostDataDto = PostUpdateDto
