package controller

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Ctrl struct{}

func (c Ctrl) ConvertCursor(cur *mongo.Cursor) []interface{} {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var arr []interface{}
	var tmp DataInput
	for cur.Next(ct) {
		cur.Decode(&tmp)
		arr = append(arr, tmp)
	}
	return arr
}

func (c Ctrl) FormatToObj(cur *mongo.SingleResult) interface{} {
	var tmp DataInput
	cur.Decode(&tmp)
	return tmp
}
