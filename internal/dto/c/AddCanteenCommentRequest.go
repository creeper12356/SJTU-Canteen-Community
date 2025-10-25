package dto

type AddCanteenCommentRequest struct {
	CanteenID uint   `binding:"required"`
	Content   string `json:"content" binding:"required"`
	Rate      uint   `json:"rate" binding:"required"`
}
