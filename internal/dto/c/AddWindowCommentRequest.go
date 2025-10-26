package dto

type AddWindowCommentRequest struct {
	WindowID   uint   `json:"window_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Rate       uint   `json:"rate" binding:"required"`
	DishID     *uint  `json:"dish_id"`
	TotalPrice *uint  `json:"total_price"`
}
