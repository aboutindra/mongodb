package http2

import (
	"encoding/json"
	"net/http"
	"time"
	"winput/controller"
)

func (i Http) InsertOneMaterial(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.DataInput
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
