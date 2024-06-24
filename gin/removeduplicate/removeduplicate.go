package removeduplicate

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
}

func RemoveDupHandle(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := remove(req.Arr1)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func remove(nums []int) []int {
	a := []int{}
	m := make(map[int]int)
	for _, value := range nums {
		m[value] = value

	}
	for key := range m {
		a = append(a, key)
	}
	return a
}
