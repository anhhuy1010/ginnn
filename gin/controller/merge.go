package controller

import (
	"fmt"
	modelMerge "gin/model/merge"
	requestMerge "gin/request/merge"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MergeArrHandle(c *gin.Context) {
	var req requestMerge.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := modelMerge.Merge(req.Arr1, req.Arr2)
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
