package model

type WindowComment struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	WindowID   uint   `gorm:"not null;index" json:"window_id"`
	Content    string `gorm:"type:text;not null" json:"content"`
	Rate       uint   `gorm:"not null" json:"rate"`
	DishID     *uint  `json:"dish_id"`
	TotalPrice *uint  `json:"total_price"`
	UserID     uint   `gorm:"not null;index" json:"user_id"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
}
