package model

import "fold/protobuf/golang/grpc/project"

type Project struct {
	Id          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Slug        string `gorm:"column:slug"`
}

func (p Project) TableName() string {
	return "projects"
}

func (p Project) ToProto() *project.CreateProjectResponse {
	resp := new(project.CreateProjectResponse)
	resp.Slug = p.Slug
	resp.Name = p.Name
	resp.Description = p.Description
	resp.Id = int32(p.Id)
	return resp
}
