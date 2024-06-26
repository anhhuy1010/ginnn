package getdetail

type (
	GetListReq struct {
		Uuid string `uri:"uuid"`
	}
	GetListRes struct {
		Uuid string `json:"uuid"`
		Task string `json:"task"`
	}
)
