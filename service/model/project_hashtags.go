package model

type ProjectHashtag struct {
	ProjectId int `gorm:"column:project_id"`
	HashtagId int `gorm:"column:hashtag_id"`
}

func (u ProjectHashtag) TableName() string {
	return "project_hashtags"
}
