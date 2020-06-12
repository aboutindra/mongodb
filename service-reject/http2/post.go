package http2

import (
	"encoding/json"
	"net/http"
	"wreject/controller"
)

func (h Http) InsertOne(res http.ResponseWriter, req *http.Request) {
	h.SetHeader(res)
	var tmpReq controller.Data
	json.NewDecoder(req.Body).Decode(&tmpReq)
	col, con := d.MakeConnection()
	err := d.InOne(col, tmpReq)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
