package dto

type ListCanteensRequest struct {
	Page     int `json:"page" binding:"required,min=0"`
	PageSize int `json:"page_size" binding:"required,min=1,max=100"`
}
