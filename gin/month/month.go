package month

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Num int `json:"num" binding:"required"`
}

func MonthHandle(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := GetDayInMonth(req.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})

}

func GetDayInMonth(month int) ([]int, error) {
	var day int
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31

	case 4, 6, 9, 11:
		day = 30

	case 2:
		day = 28

	default:
		return nil, errors.New("month is invalid")

	}
	nums := []int{}
	for i := 0; i < day; i++ {
		nums = append(nums, i+1)
	}

	return nums, nil

}
