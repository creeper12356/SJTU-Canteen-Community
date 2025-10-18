package model

type User struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string  `gorm:"type:varchar(100);not null" json:"username"`
	Avatar   *string `gorm:"type:varchar(255)" json:"avatar"`
	Email    string  `gorm:"type:varchar(100);unique;not null" json:"email"`
}
