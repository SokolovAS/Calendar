package main

import (
	"Calendar/database"
	"Calendar/entity"
	"Calendar/internal/server/http/routes"
	"log"
)

func main() {
	connection, err := database.NewGormDB()
	if err != nil {
		log.Fatal("Error db connection")
	}
	err = connection.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalln("could not migrate user model", err)
	}

	routes.BuildRouts()
}
