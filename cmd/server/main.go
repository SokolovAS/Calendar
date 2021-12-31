package main

import (
	"Calendar/database"
	"Calendar/entity"
	"Calendar/internal/server/http/routes"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
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

	//db, err := repository.InitPG()
	//if err != nil {
	//	log.Fatal("Error PG connection")
	//}
	//driver, err := postgres.WithInstance(db, &postgres.Config{})
	//m, err := migrate.NewWithDatabaseInstance(
	//	"file:///home/osoko/GolandProjects/Calendar/initdb.d/",
	//	"postgres", driver)
	//err = m.Down()
	//if err != nil {
	//	if err != migrate.ErrNoChange {
	//		log.Fatal(err)
	//	}
	//}

	routes.BuildRouts()
}
