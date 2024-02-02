package service

import (
	"context"
	"fold/protobuf/golang/grpc/hashtag"
	"fold/service/model"
	"github.com/jinzhu/gorm"
)

type hashtagService struct {
	db *gorm.DB
}

func (h hashtagService) CreateHashTag(ctx context.Context, request *hashtag.CreateHashTagRequest) (*hashtag.CreateHashTagResponse, error) {
	m := model.HashTag{Name: request.Name}
	h.db.Create(&m)
	return m.ToProto(), nil
}

func NewHashTag(db *gorm.DB) hashtag.HashtagServer {
	return &hashtagService{db: db}
}
