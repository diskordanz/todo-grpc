package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	pb "github.com/diskordanz/todo-grpc/task-service/proto"
	userService "github.com/diskordanz/todo-grpc/user-service/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"golang.org/x/net/context"
)

const (
	defaultHost = "localhost:27017"
)

var (
	srv micro.Service
)

func main() {

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	/////////////////////////////////////
	srv := micro.NewService(
		micro.Name("go.micro.srv.task"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	srv.Init()

	pb.RegisterTaskServiceHandler(srv.Server(), &Service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserService("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(ctx, &userService.Token{
			Token: token,
		})
		log.Println("Auth resp:", authResp)
		log.Println("Err:", err)
		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
