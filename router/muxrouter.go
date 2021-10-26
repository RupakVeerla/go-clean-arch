package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var m = mux.NewRouter()

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Post(url string, f func(w http.ResponseWriter, r *http.Request)) {
	m.HandleFunc(url, f).Methods("POST")
}

func (*muxRouter) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {
	m.HandleFunc(url, f).Methods("GET")
}

func (*muxRouter) Serve(port string) {
	fmt.Printf("Mux http server listening on port %s\n", port)
	http.ListenAndServe(port, m)
}
