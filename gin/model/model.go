package model

import (
	"context"
	"gin/constant"
	"gin/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoList struct {
	Uuid     string `json:"uuid" bson:"uuid"`
	Task     string `json:"task" bson:"task"`
	IsDelete int    `json:"is_delete" bson:"is_delete"`
}

func (u *TodoList) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("todos")
}

func (u *TodoList) Find(conditions map[string]interface{}) ([]*TodoList, error) {
	coll := u.Model()
	conditions["is_delete"] = constant.UNDELETE

	cursor, err := coll.Find(context.TODO(), conditions)
	if err != nil {
		return nil, err
	}

	var list []*TodoList
	for cursor.Next(context.TODO()) {
		var elem TodoList
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		list = append(list, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return list, nil
}

func (u *TodoList) FindOne(conditions map[string]interface{}) (*TodoList, error) {
	conditions["is_delete"] = constant.UNDELETE
	coll := u.Model()
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (u *TodoList) Insert() (interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return 0, err
	}

	return resp, nil
}
func (u *TodoList) Update() (int64, error) {
	coll := u.Model()

	condition := make(map[string]interface{})
	condition["uuid"] = u.Uuid

	updateStr := make(map[string]interface{})
	updateStr["$set"] = u

	resp, err := coll.UpdateOne(context.TODO(), condition, updateStr)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}
