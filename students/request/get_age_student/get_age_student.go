package getagestudent

type (
	Req struct {
		Age int `uri:"age"`
	}
	Res struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Class int    `json:"class"`
	}
)
