package main

import (
	"fmt"
	"os"
	"runtime"
	"wglobal/http2"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var htg http2.Http

func init() {
	htg = http2.Http{}
}

func main() {

	runtime.GOMAXPROCS(4)

	port := ":5554"

	handlers.AllowedHeaders([]string{"X-Requested-With"})
	handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	fmt.Println("Running in port {" + port + "}")

	r := mux.NewRouter()

	r.HandleFunc("/v1/api/record", htg.Record).Methods("POST")

	r.StrictSlash(true)

	fasthttp.ListenAndServe(port, fasthttpadaptor.NewFastHTTPHandler(r))

}
