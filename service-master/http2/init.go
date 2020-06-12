package http2

import (
	"wmaster/controller"
	"wmaster/db"
)

type Http struct{}

var d db.MongoDB
var c controller.Ctrl

func init() {
	d = db.MongoDB{"mongodb://mongo-database:27017", "WServiceMaster", "Master"}
	c = controller.Ctrl{}
}
