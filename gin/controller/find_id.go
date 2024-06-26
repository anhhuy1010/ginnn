package controller

import (
	handleFindId "gin/handle/find_id"
	requestFindId "gin/request/find_id"
	requestStudent "gin/request/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindIdHandle(c *gin.Context) {
	var req requestFindId.Req
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
	res, err := handleFindId.FindId(req.Id, Studentss)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
