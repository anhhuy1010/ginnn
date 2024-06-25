package controller

import (
	modelRemoveId "gin/model/remove_id"
	requestRemoveId "gin/request/remove_id"
	requestStudent "gin/request/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemovefindIdHandle(c *gin.Context) {
	var req requestRemoveId.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Studentss := []requestStudent.Student{
		{Name: "huy", Age: 11, Class: 12, Id: 1},
		{Name: "qưe", Age: 22, Class: 31, Id: 2},
		{Name: "aa", Age: 32, Class: 13, Id: 3},
		{Name: "voy", Age: 42, Class: 33, Id: 4},
	}
	res, err := modelRemoveId.RemoveID(req.Id, Studentss)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
