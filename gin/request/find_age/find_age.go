package findage

type Req struct {
	Age int `json:"age" binding:"required"`
}
