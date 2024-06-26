package controller

import (
	"gin/constant"
	"gin/model"
	createtodo "gin/request/create_todo"
	deletetodo "gin/request/delete_todo"
	getdetail "gin/request/get_detail"
	updatetodo "gin/request/update_todo"
	"net/http"

	getlisttodo "gin/request/get_list_todo"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController struct {
}

func (userCtl UserController) List(c *gin.Context) {
	todoModel := new(model.TodoList)
	var req getlisttodo.GetListReq
	var res []getlisttodo.GetListRes

	err := c.ShouldBindWith(&req, binding.Query)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{}
	if req.Task != nil {
		cond["task"] = *req.Task
	}
	sortby := "task"
	if req.SortBy != nil {
		sortby = *req.SortBy
	}
	sortt := 1
	if req.SortOrder != nil && *req.SortOrder != "desc" {
		sortt = -1
	}

	opts := options.Find().SetSort(bson.M{sortby: sortt})
	if req.Skip != nil {
		opts.SetSkip(*req.Skip)
	}
	if req.Limit != nil {
		opts.SetLimit(*req.Limit)
	}
	todos, err := todoModel.Find(cond, opts)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, todoo := range todos {
		resz := getlisttodo.GetListRes{
			Task: todoo.Task,
			Uuid: todoo.Uuid,
		}
		res = append(res, resz)
	}
	c.JSON(http.StatusOK, res)
}

func (userCtl UserController) Detail(c *gin.Context) {
	todoModel := new(model.TodoList)
	var req getdetail.GetListReq

	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{"uuid": req.Uuid}
	todos, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resz := getlisttodo.GetListRes{
		Task: todos.Task,
		Uuid: todos.Uuid,
	}
	c.JSON(http.StatusOK, resz)
}
func (userCtl UserController) Create(c *gin.Context) {
	var req createtodo.CreateReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todoData := model.TodoList{
		Task: req.Task,
		Uuid: req.Uuid,
	}
	_, err = todoData.Insert()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"create successfull": todoData.Task})
}
func (userCtl UserController) Update(c *gin.Context) {
	var req updatetodo.UpdateUri
	var reqq updatetodo.UpdateReq
	todoModel := new(model.TodoList)
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
	todos, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reqq.Task != "" {
		todos.Task = reqq.Task
	}
	_, err = todos.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Update successfull": todos.Uuid})
}

func (userCtl UserController) Delete(c *gin.Context) {
	todoModel := new(model.TodoList)
	var req deletetodo.DeleteUri

	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cond := bson.M{"uuid": req.Uuid}
	todos, err := todoModel.FindOne(cond)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos.IsDelete = constant.DELETE
	_, err = todos.Update()
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete succesfull": todos.Uuid})
}
