package router

import "net/http"

type Router interface {
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	Get(uri string, f func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}
