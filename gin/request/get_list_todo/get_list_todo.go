package getlisttodo

type (
	GetListReq struct {
		Uuid string `form:"uuid"`
	}
	GetListRes struct {
		Uuid string `json:"uuid"`
		Task string `json:"task"`
	}
)
