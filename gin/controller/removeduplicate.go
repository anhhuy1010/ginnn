package controller

import (
	"fmt"
	modelRemoveDuplicate "gin/model/remove_duplicate"
	requestRemoveDuplicate "gin/request/remove_duplicate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveDupHandle(c *gin.Context) {
	var req requestRemoveDuplicate.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := modelRemoveDuplicate.RemoveDuplicate(req.Arr1)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
