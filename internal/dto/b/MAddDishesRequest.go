package dto

type AddDishRequest struct {
	Name        string  `json:"name" binding:"required"`
	PictureURL  *string `json:"picture_url" binding:"required,url"`
	Description *string `json:"description" binding:"required"`
}

type MAddDishesRequest struct {
	Dishes []*AddDishRequest `json:"dishes" binding:"required,dive,required"`
}
