package router

import (
	"net/http"
)

// Router Blueprint for HTTP Query
type Router interface {
	// GET Verb for HTTP QUery
	GET(uri string, f func(resp http.ResponseWriter, req *http.Request))
	POST(uri string, f func(resp http.ResponseWriter, req *http.Request))
	SERVE(port string)
}
