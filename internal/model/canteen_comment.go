package model

type CanteenComment struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CanteenID uint   `gorm:"not null;index" json:"canteen_id"`
	Content   string `gorm:"type:text;not null" json:"content"`
	Rate      uint   `gorm:"not null" json:"rate"`
	UserID    uint   `gorm:"not null;index" json:"user_id"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
}
