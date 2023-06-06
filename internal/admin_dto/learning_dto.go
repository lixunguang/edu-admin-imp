package admin_dto

type LearningRes struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"`
	Category   string `json:"category"`
}

type LearningDetailRes struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Author     string `json:"author"`
	PictureUrl string `json:"picture_url"`
	Category   string `json:"category"`

	ResourceList []ResourceRes `json:"resource_list"`
}
