package main

import (
	"fmt"
	"os"
	"runtime"
	"winput/http2"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var hti http2.Http

func init() {
	hti = http2.Http{}
}

func main() {

	port := ":8554"

	runtime.GOMAXPROCS(4)

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.StrictSlash(true)

	r.HandleFunc("/v1/api/all", hti.GetAllData).Methods("GET")
	r.HandleFunc("/v1/api/param/i/{id}", hti.GetAllWithParamId).Methods("GET")
	r.HandleFunc("/v1/api/param/n/{name}", hti.GetAllWithParamName).Methods("GET")
	r.HandleFunc("/v1/api/total", hti.GetTotal).Methods("GET")

	r.HandleFunc("/v1/api/ino", hti.InsertOneMaterial).Methods("POST")

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
