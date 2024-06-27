package controller

import (
	"fmt"
	"net/http"
	"students/constant"
	"students/model"
	createstudent "students/request/create_student"
	deletestudent "students/request/delete_student"
	getdetailstudent "students/request/get_detail_student"
	getliststudent "students/request/get_list_student"
	updatestudent "students/request/update_student"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

type StudentController struct {
}

func (studentCtl StudentController) List(c *gin.Context) {
	todoModel := new(model.Students)
	var req getliststudent.Req
	var res []getliststudent.Res
	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{}
	student, err := todoModel.Find(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, stu := range student {
		resz := getliststudent.Res{
			Name:  stu.Name,
			Age:   stu.Age,
			Class: stu.Class,
		}
		res = append(res, resz)
	}
	c.JSON(http.StatusOK, res)
}

func (studentCtl StudentController) Detail(c *gin.Context) {
	todoModel := new(model.Students)
	var req getdetailstudent.Req
	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	cond := bson.M{"uuid": req.Uuid}
	student, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resz := getdetailstudent.Res{
		Name:  student.Name,
		Age:   student.Age,
		Class: student.Class,
	}
	c.JSON(http.StatusOK, resz)
}

func (studentCtl StudentController) Create(c *gin.Context) {
	var req createstudent.Req
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	studentData := model.Students{
		Uuid:  req.Uuid,
		Name:  req.Name,
		Age:   req.Age,
		Class: req.Class,
	}
	_, err = studentData.Insert()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"create successfull": req.Uuid})
}
func (studentCtl StudentController) Update(c *gin.Context) {
	var req updatestudent.ReqUri
	var reqq updatestudent.Req
	todoModel := new(model.Students)
	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.ShouldBindJSON(&reqq)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{"uuid": req.Uuid}
	student, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reqq.Name != "" {
		student.Name = reqq.Name
	}
	if reqq.Age != nil {
		student.Age = *reqq.Age
	}
	if reqq.Class != nil {
		student.Class = *reqq.Class
	}
	_, err = student.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Update successfull": student.Uuid})
}

func (studentCtl StudentController) Delete(c *gin.Context) {
	todoModel := new(model.Students)
	var req deletestudent.Req

	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{"uuid": req.Uuid}
	student, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student.IsDelete = constant.DELETE
	_, err = student.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete succesfull": student.Uuid})
}
