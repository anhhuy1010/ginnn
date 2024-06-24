package squares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
}

func SquaresHandle(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := squares(req.Arr1)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func squares(a []int) []int {

	new := make([]int, len(a))
	for i, val := range a {
		new[i] = val * val
	}
	n := len(new)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if new[j] > new[j+1] {
				new[j], new[j+1] = new[j+1], new[j]
			}
		}
	}
	return new
}
