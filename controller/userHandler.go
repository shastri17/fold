package controller

import (
	"context"
	userApi "fold/protobuf/golang/api/user"
	user "fold/protobuf/golang/grpc/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"strconv"
)

type UserHandler interface {
	GetUserProject(c echo.Context) error
	CreateUser(c echo.Context) error
	LinkProject(c echo.Context) error
}

type userProject struct {
	userProjectClient user.UserClient
}

func (u *userProject) LinkProject(c echo.Context) error {
	userid := c.Param("userid")
	var req userApi.LinkProjectRequest
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = protojson.UnmarshalOptions{}.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	useridint, err := strconv.Atoi(userid)
	if err != nil {
		return err
	}
	resp, err := u.userProjectClient.LinkProject(context.Background(), &user.LinkProjectRequest{
		UserId:     int32(useridint),
		ProjectIds: req.ProjectIds,
	})
	if err != nil {
		return err
	}
	return c.JSON(200, resp)
}

func (u *userProject) CreateUser(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var req userApi.CreateUserRequest
	err = protojson.UnmarshalOptions{}.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	resp, err := u.userProjectClient.CreateUser(context.Background(), &user.CreateUserRequest{Name: req.Name})
	if err != nil {
		return err
	}
	return c.JSON(200, resp)
}

func NewUserProjectController(userClient user.UserClient) UserHandler {
	return &userProject{userProjectClient: userClient}
}

func (u *userProject) GetUserProject(c echo.Context) error {
	username := c.QueryParams().Get("username")
	userid := c.QueryParams().Get("userid")
	hashtag := c.QueryParams().Get("hashtag")
	slug := c.QueryParams().Get("slug")
	desc := c.QueryParams().Get("desc")
	field := ""
	value := ""
	if username != "" {
		field = "username"
		value = username
	} else if userid != "" {
		field = "userid"
		value = userid
	} else if hashtag != "" {
		field = "hashtag"
		value = hashtag
	} else if slug != "" {
		field = "slug"
		value = slug
	} else if desc != "" {
		field = "desc"
		value = desc
	}

	resp, err := u.userProjectClient.GetUserProject(context.Background(), &user.GetUserProjectRequest{
		Field: field,
		Value: value,
	})
	if err != nil {
		return err
	}
	return c.JSON(200, resp)
}
