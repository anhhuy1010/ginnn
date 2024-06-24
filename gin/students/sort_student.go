package students

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeSorte struct {
	SortBy    string `json:"sortby" binding:"required"`
	SortOrder string `json:"sortorder" binding:"required"`
}

func SortStuHandle(c *gin.Context) {
	var req TypeSorte
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
	res, err := sort(Studentss, req.SortBy, req.SortOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func sort(a []Student, sortby string, sortorder string) ([]Student, error) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sortby == "id" && sortorder == "up" {
				if a[j].Id > a[j+1].Id {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "name" && sortorder == "up" {
				if a[j].Name > a[j+1].Name {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "age" && sortorder == "up" {
				if a[j].Age > a[j+1].Age {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "class" && sortorder == "up" {
				if a[j].Class > a[j+1].Class {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "id" && sortorder == "down" {
				if a[j].Id < a[j+1].Id {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "name" && sortorder == "down" {
				if a[j].Name < a[j+1].Name {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "age" && sortorder == "down" {
				if a[j].Age < a[j+1].Age {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else if sortby == "class" && sortorder == "down" {
				if a[j].Class < a[j+1].Class {
					a[j], a[j+1] = a[j+1], a[j]
				}
			} else {
				return nil, errors.New("sortby or sortorder incorrect")
			}
		}
	}
	return a, nil
}
