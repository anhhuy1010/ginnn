package request

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
}
