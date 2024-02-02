package main

import (
	"context"
	"fold/controller"
	"fold/protobuf/golang/grpc/hashtag"
	"fold/protobuf/golang/grpc/project"
	user "fold/protobuf/golang/grpc/user"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	ctx := context.Background()
	userServiceUrl := "grpcservice:8001"
	conn, err := grpc.DialContext(ctx, userServiceUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("error in initializing the grpc")
	}
	userClient := user.NewUserClient(conn)
	projectClient := project.NewProjectClient(conn)
	hashtagClient := hashtag.NewHashtagClient(conn)
	userhandler := controller.NewUserProjectController(userClient)
	projecthandler := controller.NewProjectHandler(projectClient)
	hashtagHandler := controller.NewHashtagHandler(hashtagClient)
	e.POST("v1/user", userhandler.CreateUser)
	e.PUT("v1/user/:userid/linkproject", userhandler.LinkProject)
	e.GET("/v1/userproject", userhandler.GetUserProject)
	e.POST("/v1/project", projecthandler.CreateProject)
	e.POST("v1/hashtag", hashtagHandler.CreateProject)
	e.PUT("v1/project/:projectid/linkhashtag", projecthandler.LinkHashTags)
	e.Logger.Fatal(e.Start(":8000"))
}
