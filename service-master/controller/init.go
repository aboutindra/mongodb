package controller

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Ctrl struct{}

func (c Ctrl) ConvertCursor(cur *mongo.Cursor) []interface{} {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var arr []interface{}
	var tmp DataMaster
	for cur.Next(ct) {
		cur.Decode(&tmp)
		arr = append(arr, tmp)
	}
	return arr
}

func (c Ctrl) FormatToObj(cur *mongo.SingleResult) interface{} {
	var tmp DataMaster
	cur.Decode(&tmp)
	return tmp
}

func (c Ctrl) FormatUpdateStok(payload int64) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"stok", payload},
		}},
	}
	return tmp
}

func (c Ctrl) FormatUpdateName(payload string) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"name", payload},
		}},
	}
	return tmp
}

func (c Ctrl) FormatUpdateModel(payload string) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"model", payload},
		}},
	}
	return tmp
}

func (c Ctrl) FormatUpdateTipe(payload string) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"tipe", payload},
		}},
	}
	return tmp
}

func (c Ctrl) FormatUpdateSub(payload []DataSub) primitive.D {
	tmp := bson.D{
		{"$set", bson.D{
			{"sub", payload},
		}},
	}
	return tmp
}
