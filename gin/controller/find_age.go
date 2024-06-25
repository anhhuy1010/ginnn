package controller

import (
	modelFindAge "gin/model/find_age"
	requestFindAge "gin/request/find_age"
	requestStudent "gin/request/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAgeHandle(c *gin.Context) {
	var req requestFindAge.Req
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
	res := modelFindAge.FindAge(req.Age, Studentss)

	c.JSON(http.StatusOK, gin.H{"result": res})
}
