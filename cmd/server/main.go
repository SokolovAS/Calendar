package main

import (
	"Calendar/database"
	"Calendar/entity"
	"Calendar/internal/server/http"
	"log"
)

func main() {
	conn := database.NewGormDB()
	err := conn.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalln("could not migrate user model", err)
	}

	http.BuildRouts()
}
