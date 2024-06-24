package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Findage struct {
	Age int `json:"age" binding:"required"`
}

func FindAgeHandle(c *gin.Context) {
	var req Findage
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
	res := findAge(req.Age, Studentss)

	c.JSON(http.StatusOK, gin.H{"result": res})
}

func findAge(age int, b []Student) []Student {
	agee := []Student{}
	for _, value := range b {
		if age < value.Age {
			agee = append(agee, value)
		}
	}
	return agee
}
