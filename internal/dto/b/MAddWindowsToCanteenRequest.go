package dto

type AddWindowRequest struct {
	Name        string  `json:"name" binding:"required"`
	PictureURL  *string `json:"picture_url"`
	Description *string `json:"description"`
}

type MAddWindowsToCanteenRequest struct {
	CanteenID uint               `binding:"required"`
	Windows   []AddWindowRequest `json:"windows" binding:"required,dive,required"`
}
