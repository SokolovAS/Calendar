package routes

import (
	"Calendar/internal/middleware"
	"Calendar/internal/repository"
	http2 "Calendar/internal/server/http"
	"Calendar/internal/services/calendar"
	"fmt"
	"log"
	"net/http"
)

const port string = ":8000"

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type router struct {
	mux   Router
	auth  *http2.AuthHandler
	event *http2.EventHandler
}

func (r *router) Init() {
	r.mux.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Up and running...")
		if err != nil {
			log.Fatalln("Error!")
		}
	})
	r.mux.GET("/events", r.event.GetAll)
	r.mux.GET("/event", r.event.GetOne)
	r.mux.POST("/event", r.event.Add)
	r.mux.PUT("/event", r.event.Update)
	r.mux.DELETE("/event", r.event.Delete)

	r.mux.POST("/signup", r.auth.Signup)
	r.mux.POST("/login", r.auth.Login)

	r.mux.SERVE(port)
}

func BuildRouts() {
	r := repository.NewSqliteRepo()
	aS := calendar.NewAuthService()
	uS := calendar.NewUserService(r)
	eS := calendar.NewEventService(r)
	mid := middleware.NewMiddleware(aS)

	aH := http2.NewAuthHandler(aS, uS)
	eH := http2.NewEventHandler(eS, uS)

	router := router{
		mux:   NewMuxRouter(mid),
		auth:  aH,
		event: eH,
	}
	router.Init()
}
