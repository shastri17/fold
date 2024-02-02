package controller

import (
	"context"
	"fold/protobuf/golang"
	projectApi "fold/protobuf/golang/api/project"
	"fold/protobuf/golang/grpc/project"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"strconv"
)

type ProjectHandler interface {
	CreateProject(c echo.Context) error
	LinkHashTags(c echo.Context) error
}

type projectHandler struct {
	projectService project.ProjectClient
}

func (p projectHandler) LinkHashTags(c echo.Context) error {
	projectid := c.Param("projectid")
	var req projectApi.LinkHashTagRequest
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = protojson.UnmarshalOptions{}.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	prjidint, err := strconv.Atoi(projectid)
	if err != nil {
		return err
	}
	resp, err := p.projectService.LinkHashtag(context.Background(), &project.LinkHashtagRequest{
		ProjectId:  int32(prjidint),
		HashtagIds: req.HashtagIds,
	})
	if err != nil {
		return err
	}
	return c.JSON(200, resp)
}

func (p projectHandler) CreateProject(c echo.Context) error {

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var req projectApi.CreateProjectRequest
	err = protojson.UnmarshalOptions{}.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	resp, err := p.projectService.CreateProject(context.Background(), &project.CreateProjectRequest{Name: req.Name, Description: req.Description})
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

func NewProjectHandler(projClient project.ProjectClient) ProjectHandler {
	return &projectHandler{projectService: projClient}
}
