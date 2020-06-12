package http2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wmaster/controller"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h Http) PutOneStokMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateStok
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateStok(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneNameMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateName
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateName(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneModelMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateModel
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateModel(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneTipeMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateTipe
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateTipe(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneSubMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateSub
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateSub(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutStokWithArr(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateArr
	json.NewDecoder(req.Body).Decode(&tmp)

	fmt.Println(tmp)

	leng := len(tmp.Data)
	//fmt.Println(leng)
	i := 0

	var tmpId controller.ItemId
	var prim primitive.D
	var tmpMaster controller.DataMaster
	var single *mongo.SingleResult
	var obj interface{}
	var err error

	for i < leng {
		tmpId.Data = tmp.Data[i].IdSub
		//fmt.Println(tmpId)
		single = d.GetOneWithParam(col, tmpId)
		obj = c.FormatToObj(single)
		//fmt.Println(obj)
		tmpMaster = obj.(controller.DataMaster)
		prim = c.FormatUpdateStok(tmpMaster.Stok - tmp.Data[i].Qty)
		err = d.UpdOne(col, tmpId, prim)

		//fmt.Println(err)

		if err != nil {
			bol.Res = false
			break
		}

		i++

	}

	if i == leng {
		bol.Res = true
	}

	d.Dis(con)

	json.NewEncoder(res).Encode(bol)

}
