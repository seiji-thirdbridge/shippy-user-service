// shippy-user-service/main.go
package main

import (
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/seiji-thirdbridge/shippy-user-service/proto/user"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Panicf("Could not connect to the database: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
