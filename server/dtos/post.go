package dtos

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
