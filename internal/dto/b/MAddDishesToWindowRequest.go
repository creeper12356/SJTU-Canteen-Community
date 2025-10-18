package dto

type DishMetadata struct {
	ID    uint  `json:"id" binding:"required"`
	Price *uint `json:"price"`
}
type MAddDishesToWindowRequest struct {
	WindowID uint           `json:"window_id" binding:"required"`
	Dishes   []DishMetadata `json:"dishes" binding:"required,dive,required"`
}
