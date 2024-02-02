package model

import "fold/protobuf/golang/grpc/user"

type User struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) ToProto() *user.CreateUserResponse {
	resp := new(user.CreateUserResponse)
	resp.Id = int32(u.Id)
	resp.Name = u.Name
	return resp
}
