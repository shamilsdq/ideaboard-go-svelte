package dtos

type PostCreateDto struct {
	SectionId int    `json:"sectionId"`
	Content   string `json:"content"`
}

type PostUpdateDto struct {
	Id        int    `json:"id"`
	SectionId int    `json:"sectionId"`
	Content   string `json:"content"`
}

type PostDeleteDto struct {
	Id int `json:"id"`
}
