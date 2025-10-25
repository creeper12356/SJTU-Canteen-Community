package dto

import "SJTU-Canteen-Community/internal/model"

type ListWindowsOfCanteenResponse struct {
	Windows []model.Window `json:"windows"`
	Total   int64          `json:"total"`
}
