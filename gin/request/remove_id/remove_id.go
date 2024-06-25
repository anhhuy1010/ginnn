package removeid

type Req struct {
	Id int `json:"id" binding:"required"`
}
