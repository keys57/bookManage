package model

type BookUser struct {
	UserID int64 `gorm:"primaryKey"`
	BookId int64 `gorm:"primaryKey"`
}
