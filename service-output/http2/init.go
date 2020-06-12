package http2

import (
	"woutput/controller"
	"woutput/db"
)

type Http struct{}

var c controller.Ctrl
var d db.MongoDB

var bol controller.ResBool

func init() {
	d = db.MongoDB{"mongodb://mongo-database:27017", "WServiceOutput", "Output"}
	c = controller.Ctrl{}
}
