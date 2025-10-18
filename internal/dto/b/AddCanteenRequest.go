package dto

type AddCanteenRequest struct {
	Name        string  `json:"name" binding:"required"`
	PictureURL  *string `json:"picture_url"`
	Description *string `json:"description"`
}
