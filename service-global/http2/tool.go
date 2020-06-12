package http2

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"wglobal/controller"
)

func (h Http) Record(res http.ResponseWriter, req *http.Request) {
	h.SetHeader(res)
	var tmpr controller.DataRequest
	json.NewDecoder(req.Body).Decode(&tmpr)

	//fmt.Println(tmpr)

	//preprare data update sub[]
	//preprae data Output
	var tmpo controller.DataOutput
	var tmpSub controller.DataSub
	var tmpArrSub controller.DataArrSub
	var tmparro controller.DataOutArr
	o := 0
	leng := len(tmpr.Data.Sub)
	for o < leng {
		tmpSub.IdSub = tmpr.Data.Sub[o].IdSub
		tmpo.IdMate = tmpr.Data.Sub[o].IdSub
		tmpSub.Name = tmpr.Data.Sub[o].Name
		tmpo.Name = tmpr.Data.Sub[o].Name
		tmpSub.Qty = tmpr.Data.Sub[o].Qty
		tmpo.Qty = tmpr.Data.Sub[o].Qty
		tmparro.Data = append(tmparro.Data, tmpo)
		tmpArrSub.Data = append(tmpArrSub.Data, tmpSub)
		o++
	}

	var wg sync.WaitGroup

	if tmpr.Sta == true {
		//prepare data Input
		var tmpi controller.DataInput
		tmpi.IdMate = tmpr.Data.IdMate
		tmpi.Name = tmpr.Data.Name
		tmpi.Qty = 1
		tmpi.Tgl = time.Now()

		//prepare data update add stok
		var tmpSubUpdate controller.FormatUpdateStok
		var tmpFind controller.ItemId
		var tmpStok controller.ItemStok

		tmpFind.Data = tmpr.Data.IdMate
		tmpStok.Data = tmpr.Data.Stok + 1
		tmpSubUpdate.Find = tmpFind
		tmpSubUpdate.Set = tmpStok

		wg.Add(4)

		go h.InputAction(tmpi, &wg)

		go h.OutputAction(tmparro, &wg)

		go h.MasterPlusAction(tmpSubUpdate, &wg)

		go h.MasterMinusAction(tmpArrSub, &wg)

	} else {
		//preprae data Reject
		var tmpre controller.DataReject
		tmpre.IdMate = tmpr.Data.IdMate
		tmpre.Name = tmpr.Data.Name
		tmpre.Model = tmpr.Data.Model
		tmpre.Tipe = tmpr.Data.Tipe

		wg.Add(3)

		go h.OutputAction(tmparro, &wg)

		go h.InputReject(tmpre, &wg)

		go h.MasterMinusAction(tmpArrSub, &wg)

	}

	wg.Wait()

	var tmpBool controller.ResBool

	tmpBool.Res = true
	json.NewEncoder(res).Encode(tmpBool)

}
