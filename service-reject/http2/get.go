package http2

import (
	"encoding/json"
	"net/http"
	"wreject/controller"
)

func (h Http) GelAll(res http.ResponseWriter, req *http.Request) {
	h.SetHeader(res)
	col, con := d.MakeConnection()
	cur, err := d.GetAll(col)
	tmp := c.CursorToObject(cur)
	if err != nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(tmp)
}

func (h Http) GetTotal(res http.ResponseWriter, req *http.Request) {
	h.SetHeader(res)
	col, con := d.MakeConnection()
	tmp, err := d.GetLen(col)

	if err != nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)

		return
	}

	var obj controller.ResInt
	obj.Res = tmp
	d.Dis(con)
	json.NewEncoder(res).Encode(obj)
}
