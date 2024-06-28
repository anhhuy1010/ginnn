package createstudent

type (
	Req struct {
		Uuid  string `json:"uuid" binding:"required"`
		Name  string `json:"name" binding:"required"`
		Class int    `json:"class" binding:"required"`
		Age   int    `json:"age" binding:"required"`
	}
	Res struct {
		Uuid string `json:"uuid"`
	}
)
