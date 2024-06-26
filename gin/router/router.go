package router

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	userCtr := new(controller.UserController)
	r := gin.Default()
	r.POST("/merge-arr", controller.MergeArrHandle)
	r.POST("/remove-duplicate", controller.RemoveDupHandle)
	r.POST("/reverse-string", controller.ReverseHandle)
	r.POST("/month", controller.MonthHandle)
	r.POST("/remove-element", controller.RemoveEleHandle)
	r.POST("/square", controller.SquaresHandle)

	r.POST("/add-controller", controller.AddStuHandle)
	r.POST("/sort", controller.SortStuHandle)
	r.POST("/find-id", controller.FindIdHandle)
	r.POST("/find-age", controller.FindAgeHandle)
	r.POST("/remove-id", controller.RemovefindIdHandle)
	r.POST("/update-controller", controller.UpdateHandle)
	r.GET("/get-list", userCtr.List)
	r.GET("/get-detail/:uuid", userCtr.Detail)
	r.POST("/create-todo", userCtr.Create)
	r.PUT("/update-todo/:uuid", userCtr.Update)
	r.DELETE("/delete-todo/:uuid", userCtr.Delete)

	return r
}
