package http2

import (
	"encoding/json"
	"net/http"
	"woutput/controller"

	"github.com/gorilla/mux"
)

func (i Http) GetAllData(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	cur, err := d.GetAll(col)
	tmp := c.ConvertCursor(cur)
	if err != nil || tmp == nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(tmp)
}

func (i Http) GetAllWithParamName(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	tmp := mux.Vars(req)["name"]
	obj := controller.ItemName{tmp}
	cur, er := d.GetAllWithParam(col, obj)
	arr := c.ConvertCursor(cur)
	if arr == nil || er != nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(arr)
}

func (i Http) GetAllWithParamId(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	tmp := mux.Vars(req)["id"]
	obj := controller.ItemId{tmp}
	cur := d.GetOneWithParam(col, obj)
	has := c.FormatToObj(cur)
	if has == nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(has)
}

func (i Http) GetTotal(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	tot, err := d.GetLen(col)
	if err != nil {
		bol.Res = false
		json.NewEncoder(res).Encode(bol)
		return
	}
	obj := struct {
		Res int64 `json:"res"`
	}{tot}
	d.Dis(con)
	json.NewEncoder(res).Encode(obj)
}
