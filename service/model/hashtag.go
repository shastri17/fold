package model

import (
	"fold/protobuf/golang/grpc/hashtag"
)

type HashTag struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (u HashTag) TableName() string {
	return "hashtags"
}

func (u HashTag) ToProto() *hashtag.CreateHashTagResponse {
	resp := new(hashtag.CreateHashTagResponse)
	resp.Id = int32(u.Id)
	resp.Name = u.Name
	return resp
}
