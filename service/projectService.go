package service

import (
	"context"
	"fmt"
	"fold/protobuf/golang/grpc/project"
	"fold/service/model"
	"github.com/jinzhu/gorm"
	"regexp"
	"strings"
)

type projectService struct {
	db *gorm.DB
}

func (p projectService) LinkHashtag(ctx context.Context, request *project.LinkHashtagRequest) (*project.LinkHashtagResponse, error) {
	resp := new(project.LinkHashtagResponse)
	for _, v := range request.HashtagIds {
		m := model.ProjectHashtag{
			ProjectId: int(request.ProjectId),
			HashtagId: int(v),
		}
		p.db.Create(&m)
	}
	return resp, nil
}

func (p projectService) CreateProject(ctx context.Context, request *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	count := 0
	p.db.Table("projects").Where(&model.Project{Name: request.Name}).Count(&count)
	slug := createSlug(request.Name)
	if count > 0 {
		slug = slug + "-" + fmt.Sprint(count)
	}
	m := model.Project{
		Name:        request.Name,
		Description: request.Description,
		Slug:        slug,
	}
	p.db.Create(&m)
	return m.ToProto(), nil
}

func NewProject(db *gorm.DB) project.ProjectServer {
	return &projectService{db: db}
}

func createSlug(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	processedString := reg.ReplaceAllString(input, " ")
	processedString = strings.TrimSpace(processedString)

	slug := strings.ReplaceAll(processedString, " ", "-")
	slug = strings.ToLower(slug)

	return slug
}
