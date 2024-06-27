package updatestudent

type (
	ReqUri struct {
		Uuid string `uri:"uuid"`
	}
	Req struct {
		Name  string `json:"name"`
		Age   *int   `json:"age" `
		Class *int   `json:"class"`
	}
)
