package dto

import "SJTU-Canteen-Community/internal/model"

type ListCanteensResponse struct {
	Canteens []model.Canteen `json:"canteens"`
	Total    int64           `json:"total"`
}
