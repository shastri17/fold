package controller

import (
	"context"
	"fold/protobuf/golang"
	hashtagApi "fold/protobuf/golang/api/hashtag"
	"fold/protobuf/golang/grpc/hashtag"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
)

type HashtagHandler interface {
	CreateProject(c echo.Context) error
}

type hashtagHandler struct {
	hashClient hashtag.HashtagClient
}

func (h hashtagHandler) CreateProject(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var req hashtagApi.CreateUserRequest
	err = protojson.UnmarshalOptions{}.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	resp, err := h.hashClient.CreateHashTag(context.Background(), &hashtag.CreateHashTagRequest{Name: req.Name})
	if err != nil {
		return c.JSON(500, Response{
			ErrorCode: errCodeMap[golang.ErrorCode_EC_INTERNAL_SERVER_ERROR],
			Data:      nil,
		})

	}
	return c.JSON(500, Response{
		Data: resp,
	})
}

func NewHashtagHandler(hashClient hashtag.HashtagClient) HashtagHandler {
	return &hashtagHandler{hashClient: hashClient}
}
