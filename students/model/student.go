package model

import (
	"context"
	"students/constant"
	"students/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type Students struct {
	Uuid     string `json:"uuid" bson:"uuid"`
	Name     string `json:"name" bson:"name"`
	Age      int    `json:"age" bson:"age"`
	Class    int    `json:"class" bson:"class"`
	IsDelete int    `json:"is_delete" bson:"is_delete"`
}

func (u *Students) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("student")
}

func (u *Students) Find(conditions map[string]interface{}) ([]*Students, error) {
	conditions["is_delete"] = constant.UNDELETE
	coll := u.Model()
	cursor, err := coll.Find(context.TODO(), conditions)
	if err != nil {
		return nil, err
	}

	var list []*Students
	for cursor.Next(context.TODO()) {
		var elem Students
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

func (u *Students) FindOne(conditions map[string]interface{}) (*Students, error) {
	conditions["is_delete"] = constant.UNDELETE
	coll := u.Model()
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (u *Students) Insert() (interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return 0, err
	}

	return resp, nil
}
func (u *Students) Update() (int64, error) {
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
