package controller

import (
	"fmt"
	handleUpdateStudent "gin/handle/update_student"
	requestStudent "gin/request/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateHandle(c *gin.Context) {
	var req requestStudent.Student
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
	res, err := handleUpdateStudent.Update(req.Id, Studentss, req)
	fmt.Println(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
