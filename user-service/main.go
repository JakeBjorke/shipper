package main

import (
	"fmt"
	"log"

	pb "github.com/jakebjorke/shipper/user-service/proto/user"
	"github.com/micro/go-micro"
)

//Authable is temp for the time being.
type Authable struct{}

func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	srv := micro.NewService(micro.Name("go.micro.srv.user"), micro.Version("latest"))
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, Authable{}})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
