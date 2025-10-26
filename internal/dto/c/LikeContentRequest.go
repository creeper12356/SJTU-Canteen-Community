package dto

type LikeContentRequest struct {
	ContentID uint `json:"content_id"`
	Weight    int  `json:"weight" binding:"required"`
}
