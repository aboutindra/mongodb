package http2

import (
	"encoding/json"
	"net/http"
	"wmaster/controller"
)

func (h Http) DelOneMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.ItemId
	json.NewDecoder(req.Body).Decode(&tmp)
	er := d.DelOne(col, tmp)
	if er != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
