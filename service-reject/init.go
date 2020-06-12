package main

import (
	"fmt"
	"os"
	"runtime"
	"wreject/http2"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var htr http2.Http

func init() {
	htr = http2.Http{}
}

func main() {

	port := ":6554"

	runtime.GOMAXPROCS(4)

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.StrictSlash(true)

	r.HandleFunc("/v1/api/all", htr.GelAll).Methods("GET")
	r.HandleFunc("/v1/api/total", htr.GetTotal).Methods("GET")

	r.HandleFunc("/v1/api/ino", htr.InsertOne).Methods("POST")

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
