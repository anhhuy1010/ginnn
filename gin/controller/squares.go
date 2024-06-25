package controller

import (
	"fmt"
	modelSquare "gin/model/square"
	requestSquares "gin/request/squares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SquaresHandle(c *gin.Context) {
	var req requestSquares.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := modelSquare.Squares(req.Arr1)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
