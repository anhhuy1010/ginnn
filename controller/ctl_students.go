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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentController struct {
}

// tìm kiếm và in hiển thị danh sách student
func (studentCtl StudentController) List(c *gin.Context) {
	studentModel := new(model.Students)
	var req getliststudent.Req
	var res []getliststudent.Res
	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{}
	// điều kiện hiển thị danh sách student theo độ tuổi
	if req.Age != nil {
		cond["age"] = bson.M{"$gt": req.Age}
	}
	// điều kiện sắp xếp danh sách tăng hoặc giảm dần theo các trường thông tin
	sortby := "name"
	if req.Sortby != nil && *req.Sortby != "" {
		sortby = *req.Sortby
	}
	sortt := 1
	if req.Sortorder != nil && *req.Sortorder == "desc" {
		sortt = -1
	}
	opts := options.Find().SetSort(bson.M{sortby: sortt})
	fmt.Println(bson.M{sortby: sortt})
	student, err := studentModel.Find(cond, opts)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, stu := range student {
		resp := getliststudent.Res{
			Name:  stu.Name,
			Age:   stu.Age,
			Class: stu.Class,
		}
		res = append(res, resp)
	}
	c.JSON(http.StatusOK, res)
}

// tìm kiếm student bằng uuid và hiển thị student
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

// tạo ra một student mới
func (studentCtl StudentController) Create(c *gin.Context) {
	var req createstudent.Req
	studentModel := new(model.Students)
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// điều kiện kiểm tra uuid đã có trong dữ liệu hay chưa
	cond := bson.M{"uuid": req.Uuid}
	_, err = studentModel.FindOne(cond)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Student already exists"})
		return
	}

	studentData := model.Students{
		Uuid:     req.Uuid,
		Name:     req.Name,
		Age:      req.Age,
		Class:    req.Class,
		IsDelete: constant.UNDELETE,
	}
	_, err = studentData.Insert()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(studentData)

	c.JSON(http.StatusOK, gin.H{"create successfull": req.Uuid})
}

// hàm cập nhật student
func (studentCtl StudentController) Update(c *gin.Context) {
	var reqUri updatestudent.ReqUri
	var req updatestudent.Req
	todoModel := new(model.Students)
	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.ShouldBindJSON(&reqUri)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{"uuid": reqUri.Uuid}
	student, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name != "" {
		student.Name = req.Name
	}
	if req.Age != nil {
		student.Age = *req.Age
	}
	if req.Class != nil {
		student.Class = *req.Class
	}
	_, err = student.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Update successfull": student.Uuid})
}

// hàm xóa một student
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
