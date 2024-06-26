package controller

import (
	"fmt"
	handleReverseString "gin/handle/reverse_string"
	requestReverseString "gin/request/reverse_string"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReverseHandle(c *gin.Context) {
	var req requestReverseString.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := handleReverseString.Reverse(req.Arr)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
