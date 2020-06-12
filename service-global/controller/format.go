package controller

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataSub struct {
	IdSub string `json:"idSub, omitempty" bson:"idSub, omitempty"`
	Name  string `json:"name, omitempty bson:"name, omitempty""`
	Qty   int64  `json:"qty, omitempty" bson:"qty, omitempty"`
}

type DataMaster struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMate  string             `json:"idMate" bson:"idMate"`
	Name    string             `json:"name, omitempty" bson:"name, omitempty"`
	Model   string             `json:"model, omitempty" bson:"model, omitempty"`
	Tipe    string             `json:"tipe, omitempty" bson:"tipe,omitempty"`
	Stok    int64              `json:"stok, omitempty" bson:"stok, omitempty"`
	Sopname int64              `json:"sopname, omitempty" bson:"sopname, omitempty"`
	Tgl     time.Time          `json:"tgl,omitempty" bson:"tgl,omitempty"`
	Sub     []DataSub          `json:"sub" bson:"sub"`
}

type DataArrSub struct {
	Data []DataSub `json:"data" bson:"data"`
}

type DataRequest struct {
	Data DataMaster `json:"data" bson:"data"`
	Sta  bool       `json:"sta" bson:"sta"`
}

type ItemId struct {
	Data string `json:"idMate, omitempty" bson:"idMate, omitempty"`
}

type ItemStok struct {
	Data int64 `json:"stok,omitempty" bson:"stok,omitempty"`
}

type FormatUpdateStok struct {
	Find ItemId   `json:"find" bson:"find"`
	Set  ItemStok `json:"set" bson:"set"`
}

type DataInput struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMate string             `json:"idMate, omitempty" bson:"idMate, omitempty"`
	Name   string             `json:"name, omitempty" bson:"name, omitempty"`
	Qty    int64              `json:"qty, omitempty" bson:"qty, omitempty"`
	Tgl    time.Time          `json:"tgl" bson:"tgl"`
}

type DataOutput struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMate string             `json:"idMate, omitempty" bson:"idMate, omitempty"`
	Name   string             `json:"name, omitempty" bson:"name, omitempty"`
	Qty    int64              `json:"qty, omitempty" bson:"qty, omitempty"`
}

type DataReject struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdMate string             `json:"idMate,omitempty" bson:"idMate, omitempty"`
	Name   string             `json:"name, omitempty" bson:"name, omitempty"`
	Model  string             `json:"model, omitempty" bson:"model, omitempty"`
	Tipe   string             `json:"tipe, omitempty" bson:"tipe,omitempty"`
}

type DataOutArr struct {
	Data []DataOutput `json:"data" bson:"data"`
}

type ResBool struct {
	Res bool `json:"res" bson:"res"`
}
