package http2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"wglobal/controller"
)

type Http struct{}

var c controller.Ctrl

func init() {
	c = controller.Ctrl{}
}

func (h Http) InputAction(payload controller.DataInput, wg *sync.WaitGroup) {

	reqb, _ := json.Marshal(payload)

	resp, _ := http.Post("http://localhost:8554/v1/api/ino", "application/json", bytes.NewBuffer(reqb))

	defer resp.Body.Close()

	defer wg.Done()

}

func (h Http) OutputAction(payload controller.DataOutArr, wg *sync.WaitGroup) {

	reqb, _ := json.Marshal(payload)

	resp, _ := http.Post("http://localhost:9554/v1/api/inm", "application/json", bytes.NewBuffer(reqb))

	defer resp.Body.Close()

	defer wg.Done()

}

func (h Http) MasterPlusAction(payload controller.FormatUpdateStok, wg *sync.WaitGroup) {

	reqb, _ := json.Marshal(payload)

	rep, _ := http.NewRequest(http.MethodPut, "http://localhost:7554/v1/api/seto/stok", bytes.NewBuffer(reqb))

	rep.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, _ := client.Do(rep)

	defer resp.Body.Close()

	defer wg.Done()

}

func (h Http) MasterMinusAction(payload controller.DataArrSub, wg *sync.WaitGroup) {

	fmt.Println(payload)

	reqb, _ := json.Marshal(payload)

	rep, _ := http.NewRequest(http.MethodPut, "http://localhost:7554/v1/api/setm/stok", bytes.NewBuffer(reqb))

	rep.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, _ := client.Do(rep)

	defer resp.Body.Close()

	defer wg.Done()

}

func (h Http) InputReject(payload controller.DataReject, wg *sync.WaitGroup) {

	reqb, _ := json.Marshal(payload)

	resp, _ := http.Post("http://localhost:6554/v1/api/ino", "application/json", bytes.NewBuffer(reqb))

	defer resp.Body.Close()

	defer wg.Done()

}
