package entities

type Section struct {
	Title   string
	PostIds []int
}

func (section *Section) AddPostId(postIdToBeAdded int) {
	for _, postId := range section.PostIds {
		if postId == postIdToBeAdded {
			return
		}
	}
	section.PostIds = append(section.PostIds, postIdToBeAdded)
}

func (section *Section) RemovePostId(postIdToBeRemoved int) {
	for idx, postId := range section.PostIds {
		if postId == postIdToBeRemoved {
			section.PostIds = append(section.PostIds[:idx], section.PostIds[idx+1:]...)
		}
	}
}
