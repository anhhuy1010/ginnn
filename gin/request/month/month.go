package month

type Req struct {
	Num int `json:"num" binding:"required"`
}
