package students

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateHandle(c *gin.Context) {
	var req Student
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
	res, err := update(req.Id, Studentss, req)
	fmt.Println(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
func update(nums int, b []Student, p Student) ([]Student, error) {
	reg := false
	for i, value := range b {
		if value.Id == nums {
			b[i].Name = p.Name
			b[i].Age = p.Age
			b[i].Class = p.Class
			reg = true
			break
		}
	}
	if !reg {
		return nil, errors.New("id does not exist")
	}

	return b, nil
}
