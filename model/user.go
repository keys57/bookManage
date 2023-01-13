package model

type User struct {
	Id       int64  `json:"id" grom:"primaryKey"`
	UserName string `json:"username" gorm:"not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Token    string `json:"token"`
}

func (User) TableName() string {
	return "user"
}
