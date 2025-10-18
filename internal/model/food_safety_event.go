package model

type FoodSafetyEvent struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string `gorm:"type:varchar(200);not null" json:"title"`
	WindowID        uint   `gorm:"not null;index" json:"window_id"`
	Content         string `gorm:"type:text;not null" json:"content"`
	WindowCommentID uint   `gorm:"not null;index" json:"window_comment_id"`
	UserID          uint   `gorm:"not null;index" json:"user_id"`
	CreatedAt       int64  `gorm:"autoCreateTime" json:"created_at"`
}
