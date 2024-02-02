package model

type UserProject struct {
	UserId    int `gorm:"column:user_id"`
	ProjectId int `gorm:"column:project_id"`
}

func (u UserProject) TableName() string {
	return "user_projects"
}
