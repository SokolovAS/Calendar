package http

import (
	middlewares "Calendar/internal/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
	muxDispatcher.Use(middlewares.Authz)
}
func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
	muxDispatcher.Use(middlewares.Authz)
}
func (m *muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
	muxDispatcher.Use(middlewares.Authz)
}
func (m *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
	muxDispatcher.Use(middlewares.Authz)
}
func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux http server is running on port %v", port)
	err := http.ListenAndServe(port, muxDispatcher)
	if err != nil {
		log.Fatal(err)
	}
}
