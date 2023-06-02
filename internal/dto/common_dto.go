package dto

//通用的输入参数，只有ID一个参数
type IDParam struct {
	ID int `json:"id" binding:"required"`
}

//通用的返回参数，只有ID一个参数
type IDRes struct {
	ID int `json:"id"`
}
