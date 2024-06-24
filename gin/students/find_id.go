package students

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Findid struct {
	Id int `json:"id" binding:"required"`
}

func FindIdHandle(c *gin.Context) {
	var req Findid
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Studentss := []Student{
		{Name: "huy", Age: 11, Class: 12, Id: 1},
		{Name: "qưe", Age: 22, Class: 31, Id: 2},
		{Name: "aa", Age: 32, Class: 13, Id: 3},
		{Name: "voy", Age: 42, Class: 33, Id: 4},
	}
	res, err := findId(req.Id, Studentss)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func findId(id int, b []Student) (Student, error) {
	for _, student := range b {
		if student.Id == id {
			return student, nil
		}
	}

	return Student{}, errors.New("ID does not exist")
}
