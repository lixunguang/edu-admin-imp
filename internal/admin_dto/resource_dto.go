package admin_dto

type ResouceIDParam struct {
	ID int `json:"id"`
}

type ResouceDelParam struct {
	ID   int `json:"id"`
	Type int `json:"type"`
}

type ResourceParam struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type UpdateRichtextParam struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
