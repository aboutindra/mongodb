package http2

import "net/http"

func (h Http) SetHeader(res http.ResponseWriter) {

	res.Header().Add("content-type", "application/json")

}
