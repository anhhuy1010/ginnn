package controller

import (
	handleSort "gin/handle/sort_student"
	requestSortStudent "gin/request/sort_student"
	requestStudent "gin/request/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SortStuHandle(c *gin.Context) {
	var req requestSortStudent.Req
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
	res, err := handleSort.Sort(Studentss, req.SortBy, req.SortOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
