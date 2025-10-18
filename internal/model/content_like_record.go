package model

type ContentLikeRecord struct {
	ContentType uint  `gorm:"not null;index:idx_content_type_content_id,unique" json:"content_type"`
	ContentID   uint  `gorm:"not null;index:idx_content_type_content_id,unique" json:"content_id"`
	Weight      int   `gorm:"not null" json:"weight"`
	UserID      uint  `gorm:"not null;index" json:"user_id"`
	CreatedAt   int64 `gorm:"autoCreateTime" json:"created_at"`
}
