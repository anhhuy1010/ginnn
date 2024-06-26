package createtodo

type (
	CreateReq struct {
		Uuid     string `json:"uuid" binding:"required"`
		Task     string `json:"task" binding:"required"`
		IsDelete int    `json:"is_delete"`
	}
	GetListRes struct {
		Uuid     string `json:"uuid"`
		Task     string `json:"task"`
		IsDelete int    `json:"is_delete"`
	}
)
