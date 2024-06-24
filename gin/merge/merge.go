package merge

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
	Arr2 []int `json:"arr2" binding:"required"`
}

func MergeArrHandle(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := Merge(req.Arr1, req.Arr2)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func Merge(arr1 []int, arr2 []int) []int {
	arr := []int{}
	arr = append(arr1, arr2...)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
