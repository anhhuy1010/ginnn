package students

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Removeid struct {
	Id int `json:"id" binding:"required"`
}

func RemovefindIdHandle(c *gin.Context) {
	var req Removeid
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
	res, err := remove(req.Id, Studentss)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func remove(id int, b []Student) ([]Student, error) {
	removee := []Student{}
	r := false
	for _, value := range b {
		if id == value.Id {
			r = true
		} else {
			removee = append(removee, value)
		}
	}
	if !r {
		return removee, errors.New("id does not exist")
	}
	return removee, nil
}
