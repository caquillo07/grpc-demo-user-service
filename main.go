package main

import (
	"fmt"
	pb "github.com/caquillo07/grpc-demo-user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// Creates a database connection and handles
	// closing it before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("could not connect to DB: %v\n", err)
	}

	fmt.Printf("connected to postgres %+v\n", db, err)

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("user service is listening")
}
