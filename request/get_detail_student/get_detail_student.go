package getdetailstudent

type (
	Req struct {
		Uuid string `uri:"uuid"`
	}
	Res struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Class int    `json:"class"`
	}
)
