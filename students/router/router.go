package router

import (
	"students/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	studentCtr := new(controller.StudentController)
	r := gin.Default()
	r.GET("/get-list", studentCtr.List)
	r.GET("/get-detail/:uuid", studentCtr.Detail)
	r.POST("/create-student", studentCtr.Create)
	r.PUT("/update-student/:uuid", studentCtr.Update)
	r.DELETE("/delete-student/:uuid", studentCtr.Delete)
	return r

}
