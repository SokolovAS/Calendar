package main

import (
	"Calendar/controller"
	"Calendar/entity"
	database "Calendar/initdb.d"
	router "Calendar/internal/server/http"
	"fmt"
	"log"
	"net/http"
)

var (
	eventController controller.EventController = controller.NewEventController()
	authController  controller.AuthController  = controller.NewAuthController()
	httpRouter      router.Router              = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	err = database.GlobalDB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalln("could not migrate user model", err)
	}

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
