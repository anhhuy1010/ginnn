package reversestring

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Arr string `json:"arr1" binding:"required"`
}

func ReverseHandle(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := reverse(req.Arr)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func reverse(arr string) string {
	var new string
	for i := len(arr) - 1; i >= 0; i-- {
		new = new + string(arr[i])
	}
	return new
}
