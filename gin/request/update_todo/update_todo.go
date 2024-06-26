package updatetodo

type (
	UpdateUri struct {
		Uuid string `uri:"uuid"`
	}
	UpdateReq struct {
		Task string `json:"task" binding:"required"`
	}
)
