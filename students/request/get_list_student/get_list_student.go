package getliststudent

type (
	Req struct {
		Name  string `form:"name"`
		Age   int    `form:"age"`
		Class int    `form:"class"`
	}
	Res struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Class int    `json:"class"`
	}
)
