package main

import (
	"Calendar/database"
	"Calendar/entity"
	"Calendar/interceptors"
	"Calendar/internal/repository"
	"Calendar/internal/server/grps"
	"Calendar/internal/services/calendar"
	"Calendar/pb"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
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

	db, err := repository.InitPG()
	if err != nil {
		log.Fatal("Error PG connection")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///home/osoko/GolandProjects/Calendar/initdb.d/",
		"postgres", driver)
	err = m.Down()
	if err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	//routes.BuildRouts()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authMD := interceptors.AuthMD{}
	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.ChainUnaryInterceptor(authMD.UnaryInterceptor()))

	grpcServer := grpc.NewServer(opts...)

	r := repository.NewRepoPG()
	es := calendar.NewEventService(r)
	s := grps.NewGRPCServer(es)
	// registering specific handlers for this server
	pb.RegisterEventServiceServer(grpcServer, s)
	log.Println("starting server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
