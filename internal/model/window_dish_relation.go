package model

type WindowDishRelation struct {
	ID       uint `gorm:"primaryKey;autoIncrement" json:"id"`
	WindowID uint `gorm:"not null;index" json:"window_id"`
	DishID   uint `gorm:"not null" json:"dish_id"`
	Price    uint `json:"price"`
}
