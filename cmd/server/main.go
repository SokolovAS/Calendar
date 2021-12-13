package main

import (
	"Calendar/controller"
	router "Calendar/internal/server/http"
	"fmt"
	"log"
	"net/http"
)

var (
	eventController controller.EventController = controller.NewEventController()
	httpRouter      router.Router              = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

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

	httpRouter.SERVE(port)
}
