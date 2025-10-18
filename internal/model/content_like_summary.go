package model

type ContentLikeSummary struct {
	ContentType uint `gorm:"not null;index:idx_content_type_content_id,unique" json:"content_type"`
	ContentID   uint `gorm:"not null;index:idx_content_type_content_id,unique" json:"content_id"`
	LikeCount   int  `gorm:"not null" json:"like_count"`
}
