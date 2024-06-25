package merge

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
	Arr2 []int `json:"arr2" binding:"required"`
}
