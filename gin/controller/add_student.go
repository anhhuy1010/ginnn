package controller

import (
	"fmt"
	modelAddStudent "gin/model/add_student"
	requestStudent "gin/request/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStuHandle(c *gin.Context) {
	var req requestStudent.Student

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Studentss := []requestStudent.Student{
		{Name: "hy", Age: 11, Class: 12, Id: 1},
		{Name: "py", Age: 22, Class: 31, Id: 2},
		{Name: "uy", Age: 32, Class: 13, Id: 3},
		{Name: "vit", Age: 42, Class: 33, Id: 4},
	}

	res, err := modelAddStudent.Add_Student(Studentss, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
