package removeelement

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Arr1 []int `json:"arr1" binding:"required"`
	Num  int   `json:"num" binding:"required"`
}

func RemoveEleHandle(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, ress := removeElement(req.Arr1, req.Num)
	fmt.Println(res, ress)
	c.JSON(http.StatusOK, gin.H{"val": res, "result": ress})

}

func removeElement(nums []int, value int) (int, []int) {

	a := []int{}
	for _, val := range nums {
		if val != value {
			a = append(a, val)
		}
	}
	n := len(a)
	return n, a
}
