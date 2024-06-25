package student

type Student struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Class int    `json:"class" binding:"required"`
	Id    int    `json:"id" binding:"required"`
}
