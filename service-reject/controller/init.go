package controller

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Ctrl struct{}

func (c Ctrl) CursorToObject(payload *mongo.Cursor) []interface{} {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var arr []interface{}
	var tmp Data
	for payload.Next(ct) {
		payload.Decode(&tmp)
		arr = append(arr, tmp)
	}
	return arr
}
