package controller

import (
	"fmt"
	modelRemoveElement "gin/model/remove_element"
	requestRemoveElement "gin/request/removeelement"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveEleHandle(c *gin.Context) {
	var req requestRemoveElement.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, ress := modelRemoveElement.RemoveElement(req.Arr1, req.Num)
	fmt.Println(res, ress)
	c.JSON(http.StatusOK, gin.H{"val": res, "result": ress})

}
