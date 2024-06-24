package main

import (
	"gin/merge"
	"gin/month"
	"gin/removeduplicate"
	"gin/removeelement"
	"gin/reversestring"
	"gin/squares"
	"gin/students"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/merge-arr", merge.MergeArrHandle)
	r.POST("/remove-duplicate", removeduplicate.RemoveDupHandle)
	r.POST("/reverse-string", reversestring.ReverseHandle)
	r.POST("/month", month.MonthHandle)
	r.POST("/remove-element", removeelement.RemoveEleHandle)
	r.POST("/square", squares.SquaresHandle)

	r.POST("/add-student", students.AddStuHandle)
	r.POST("/sort", students.SortStuHandle)
	r.POST("/find-id", students.FindIdHandle)
	r.POST("/find-age", students.FindAgeHandle)
	r.POST("/remove-id", students.RemovefindIdHandle)
	r.POST("/update-student", students.UpdateHandle)

	return r
}
func main() {
	r := setupRouter()
	r.Run(":8088")
}
