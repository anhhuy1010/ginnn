package controller

import (
	modelMonth "gin/model/month"
	requestMonth "gin/request/month"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MonthHandle(c *gin.Context) {
	var req requestMonth.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := modelMonth.GetDayInMonth(req.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})

}
