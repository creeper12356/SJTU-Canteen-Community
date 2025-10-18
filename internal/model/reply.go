package model

type Reply struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CommentType uint   `gorm:"not null;index:idx_comment_type_comment_id" json:"comment_type"`
	CommentID   uint   `gorm:"not null;index:idx_comment_type_comment_id" json:"comment_id"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	ReplyTo     uint   `gorm:"index" json:"reply_to"`
	Content     string `gorm:"type:text;not null" json:"content"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
}
