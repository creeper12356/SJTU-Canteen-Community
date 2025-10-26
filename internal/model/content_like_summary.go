package model

import "SJTU-Canteen-Community/internal/consts"

type ContentLikeSummary struct {
	ContentType consts.ContentType `gorm:"primaryKey;not null;index:idx_content_type_content_id,unique" json:"content_type"`
	ContentID   uint               `gorm:"primaryKey;not null;index:idx_content_type_content_id,unique" json:"content_id"`
	LikeCount   int                `gorm:"not null" json:"like_count"`
}
