package http2

import (
	"encoding/json"
	"net/http"
	"woutput/controller"
)

func (o Http) InsertManyMaterial(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmpReq controller.DataOutArr
	json.NewDecoder(req.Body).Decode(&tmpReq)
	var tmpParam []interface{}
	len := len(tmpReq.Data)
	var i = 0
	for i < len {
		var tmp controller.DataOutput
		tmp = tmpReq.Data[i]
		tmpParam = append(tmpParam, tmp)
		i++
	}
	err := d.InMany(col, tmpParam)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
