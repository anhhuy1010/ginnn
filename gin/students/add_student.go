package students

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Class int    `json:"class" binding:"required"`
	Id    int    `json:"id" binding:"required"`
}

func AddStuHandle(c *gin.Context) {
	var req Student

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Studentss := []Student{
		{Name: "hy", Age: 11, Class: 12, Id: 1},
		{Name: "py", Age: 22, Class: 31, Id: 2},
		{Name: "uy", Age: 32, Class: 13, Id: 3},
		{Name: "vit", Age: 42, Class: 33, Id: 4},
	}

	res, err := add_student(Studentss, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func add_student(student []Student, person Student) ([]Student, error) {
	m := make(map[int]Student)
	for _, value := range student {
		m[value.Id] = value
	}
	_, ok := m[person.Id]
	if ok {

		return student, errors.New("ID already exists")

	} else {
		student = append(student, person)
	}

	return student, nil
}
