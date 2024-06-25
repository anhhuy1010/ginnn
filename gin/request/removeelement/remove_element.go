package removeelement

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
	Num  int   `json:"num" binding:"required"`
}
