package findid

type Req struct {
	Id int `json:"id" binding:"required"`
}
