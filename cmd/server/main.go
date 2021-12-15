package main

import (
	"Calendar/entity"
	database "Calendar/initdb.d"
	router "Calendar/internal/server/http"
	"log"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	err = database.GlobalDB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalln("could not migrate user model", err)
	}

	router.Init()
}
