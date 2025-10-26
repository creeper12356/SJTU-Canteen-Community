package repository

import (
	"SJTU-Canteen-Community/internal/consts"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type ContentLikeSummaryRepository struct {
	db *gorm.DB
}

func NewContentLikeSummaryRepository(db *gorm.DB) *ContentLikeSummaryRepository {
	return &ContentLikeSummaryRepository{db: db}
}

func (r *ContentLikeSummaryRepository) FindContentLikeSummary(contentType consts.ContentType, contentID uint, forUpdate bool) (model.ContentLikeSummary, error) {
	var summary model.ContentLikeSummary
	db := r.db
	if forUpdate {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}
	result := db.Where("content_type = ? AND content_id = ?", contentType, contentID).First(&summary)
	return summary, result.Error
}

func (r *ContentLikeSummaryRepository) UpdateOrCreateContentLikeSummary(summary *model.ContentLikeSummary) error {
	result := r.db.Save(summary)
	return result.Error
}
