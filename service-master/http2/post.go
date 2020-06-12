package http2

import (
	"encoding/json"
	"net/http"
	"time"
	"wmaster/controller"
)

func (h Http) InsertOneMaterial(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.DataMaster
	json.NewDecoder(req.Body).Decode(&tmp)
	tmp.Tgl = time.Now()
	er := d.InOne(col, tmp)
	if er != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) InsertManyMaterial(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmpReq []controller.DataMaster
	json.NewDecoder(req.Body).Decode(&tmpReq)
	var tmpParam []interface{}
	len := len(tmpReq)
	i := 0
	for i < len {
		var tmp controller.DataMaster
		tmp = tmpReq[i]
		tmp.Tgl = time.Now()
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
