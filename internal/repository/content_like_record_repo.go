package repository

import (
	"SJTU-Canteen-Community/internal/consts"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type ContentLikeRecordRepository struct {
	db *gorm.DB
}

func NewContentLikeRecordRepository(db *gorm.DB) *ContentLikeRecordRepository {
	return &ContentLikeRecordRepository{db: db}
}

func (r *ContentLikeRecordRepository) FindContentLikeRecord(userID uint, contentType consts.ContentType, contentID uint, forUpdate bool) (model.ContentLikeRecord, error) {
	var record model.ContentLikeRecord
	db := r.db
	if forUpdate {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}
	result := db.Where("user_id = ? AND content_type = ? AND content_id = ?", userID, contentType, contentID).First(&record)
	return record, result.Error
}

func (r *ContentLikeRecordRepository) UpdateOrCreateContentLikeRecord(record *model.ContentLikeRecord) error {
	result := r.db.Save(record)
	return result.Error
}

func (r *ContentLikeRecordRepository) DeleteContentLikeRecord(record *model.ContentLikeRecord) error {
	result := r.db.Delete(record)
	return result.Error
}
