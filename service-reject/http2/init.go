package http2

import (
	"net/http"
	"wreject/controller"
	"wreject/db"
)

var d db.MongoDB
var c controller.Ctrl

var bol controller.ResBool

type Http struct{}

func init() {
	d = db.MongoDB{"mongodb://mongo-database:27017", "WServiceReject", "Reject"}
	c = controller.Ctrl{}
}

func (h Http) SetHeader(payload http.ResponseWriter) {
	payload.Header().Add("content-type", "application/json")
}
