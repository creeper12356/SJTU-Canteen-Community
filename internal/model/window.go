package model

type Window struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	PictureURL  string `gorm:"type:varchar(255)" json:"picture_url"`
	Description string `gorm:"type:text" json:"description"`
	CanteenID   uint   `gorm:"not null;index" json:"canteen_id"`
}
