package http

import (
	middlewares "Calendar/internal/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http"
)

type muxRouter struct {
	dispatcher *mux.Router
	mid        middlewares.Middleware
}

func NewMuxRouter() Router {
	return &muxRouter{
		dispatcher: mux.NewRouter(),
		mid:        middlewares.NewMiddleware(),
	}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.dispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
	m.dispatcher.Use(m.mid.Authz)
}
func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.dispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
	m.dispatcher.Use(m.mid.Authz)
}
func (m *muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.dispatcher.HandleFunc(uri, f).Methods(http.MethodPut)
	m.dispatcher.Use(m.mid.Authz)
}
func (m *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.dispatcher.HandleFunc(uri, f).Methods(http.MethodDelete)
	m.dispatcher.Use(m.mid.Authz)
}
func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux http server is running on port %v", port)
	err := http.ListenAndServe(port, m.dispatcher)
	if err != nil {
		log.Fatal(err)
	}
}
