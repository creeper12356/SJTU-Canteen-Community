package model

import "SJTU-Canteen-Community/internal/consts"

type ContentLikeRecord struct {
	UserID      uint               `gorm:"primaryKey;not null;index:user_id_idx_content_type_content_id,unique" json:"user_id"`
	ContentType consts.ContentType `gorm:"primaryKey;not null;index:content_type_idx_user_id_content_id,unique" json:"content_type"`
	ContentID   uint               `gorm:"primaryKey;not null;index:content_id_idx_user_id_content_type,unique" json:"content_id"`
	Weight      int                `gorm:"not null" json:"weight"`
	CreatedAt   int64              `gorm:"autoCreateTime" json:"created_at"`
}
