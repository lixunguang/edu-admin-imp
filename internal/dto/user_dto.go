package dto

type User struct {
	LoginID        string `json:"login_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Password       string `json:"password" binding:"required"`
	OrganizationID int    `json:"organization_id"`
}

type UserRes struct {
	Name string `json:"name"`
	//Role    int    `json:"role"`
	LoginId        string `json:"login_id"`
	OrganizationID int    `json:"organization_id"`
}

type AddUserRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
