package http

import (
	"Calendar/controller"
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

var (
	eventController controller.EventController = controller.NewEventController()
	authController  controller.AuthController  = controller.NewAuthController()
	httpRouter      Router                     = NewMuxRouter()
)

func Init() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Up and running...")
		if err != nil {
			log.Fatalln("Error!")
		}
	})
	httpRouter.GET("/events", eventController.GetAll)
	httpRouter.GET("/event", eventController.GetOne)
	httpRouter.POST("/event", eventController.Add)
	httpRouter.PUT("/event", eventController.Update)
	httpRouter.DELETE("/event", eventController.Delete)

	httpRouter.POST("/signup", authController.Signup)
	httpRouter.POST("/login", authController.Login)

	httpRouter.SERVE(port)
}
