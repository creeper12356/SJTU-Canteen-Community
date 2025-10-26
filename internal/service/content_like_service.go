package service

import (
	"SJTU-Canteen-Community/internal/consts"
	"SJTU-Canteen-Community/internal/model"
	"SJTU-Canteen-Community/internal/repository"
	"fmt"

	"gorm.io/gorm"
)

type ContentLikeService struct {
	db *gorm.DB
}

func NewContentLikeService(db *gorm.DB) *ContentLikeService {
	return &ContentLikeService{db: db}
}

func (s *ContentLikeService) LikeContent(checkContentExists func(db *gorm.DB, contentID uint) (bool, error), contentType consts.ContentType, contentID uint, userID uint, weight int) error {
	if weight != 1 && weight != -1 {
		return fmt.Errorf("invalid weight: %d, weight must be either 1 or -1", weight)
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	exists, err := checkContentExists(tx, contentID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if !exists {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	contentLikeRecordRepo := repository.NewContentLikeRecordRepository(tx)
	record, err := contentLikeRecordRepo.FindContentLikeRecord(userID, contentType, contentID, true)

	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return err
	}

	likeCountDelta := 0
	if err == gorm.ErrRecordNotFound {
		record = model.ContentLikeRecord{
			UserID:      userID,
			ContentType: contentType,
			ContentID:   contentID,
			Weight:      weight,
		}
		likeCountDelta = weight
	} else {
		likeCountDelta = weight - record.Weight
		record.Weight = weight
	}

	err = contentLikeRecordRepo.UpdateOrCreateContentLikeRecord(&record)
	if err != nil {
		tx.Rollback()
		return err
	}

	if likeCountDelta != 0 {
		var summary model.ContentLikeSummary
		contentLikeSummaryRepo := repository.NewContentLikeSummaryRepository(tx)
		summary, err := contentLikeSummaryRepo.FindContentLikeSummary(contentType, contentID, true)
		if err != nil && err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return err
		}

		if err == gorm.ErrRecordNotFound {
			summary = model.ContentLikeSummary{
				ContentType: contentType,
				ContentID:   contentID,
				LikeCount:   likeCountDelta,
			}
		} else {
			summary.LikeCount += likeCountDelta
		}

		err = contentLikeSummaryRepo.UpdateOrCreateContentLikeSummary(&summary)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (s *ContentLikeService) CancelLikeContent(checkContentExists func(db *gorm.DB, contentID uint) (bool, error), contentType consts.ContentType, contentID uint, userID uint) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	exists, err := checkContentExists(tx, contentID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !exists {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	contentLikeRecordRepo := repository.NewContentLikeRecordRepository(tx)
	record, err := contentLikeRecordRepo.FindContentLikeRecord(userID, contentType, contentID, true)

	if err != nil {
		tx.Rollback()
		return err
	}

	likeCountDelta := -record.Weight
	if err := contentLikeRecordRepo.DeleteContentLikeRecord(&record); err != nil {
		tx.Rollback()
		return err
	}

	contentLikeSummaryRepo := repository.NewContentLikeSummaryRepository(tx)

	summary, err := contentLikeSummaryRepo.FindContentLikeSummary(contentType, contentID, true)
	if err != nil {
		tx.Rollback()
		return err
	}

	summary.LikeCount += likeCountDelta

	if err := contentLikeSummaryRepo.UpdateOrCreateContentLikeSummary(&summary); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func checkCanteenCommentExists(db *gorm.DB, contentID uint) (bool, error) {
	canteenCommentRepo := repository.NewCanteenCommentRepository(db)
	return canteenCommentRepo.CheckCanteenCommentExists(contentID, true)
}

func (s *ContentLikeService) LikeCanteenComment(canteenCommentID uint, userID uint, weight int) error {
	return s.LikeContent(checkCanteenCommentExists, consts.ContentType_CanteenComment, canteenCommentID, userID, weight)
}

func (s *ContentLikeService) CancelLikeCanteenComment(canteenCommentID uint, userID uint) error {
	return s.CancelLikeContent(checkCanteenCommentExists, consts.ContentType_CanteenComment, canteenCommentID, userID)
}

func checkWindowCommentExists(db *gorm.DB, contentID uint) (bool, error) {
	windowCommentRepo := repository.NewWindowCommentRepository(db)
	return windowCommentRepo.CheckWindowCommentExists(contentID, true)
}

func (s *ContentLikeService) LikeWindowComment(windowCommentID uint, userID uint, weight int) error {
	return s.LikeContent(checkWindowCommentExists, consts.ContentType_WindowComment, windowCommentID, userID, weight)
}

func (s *ContentLikeService) CancelLikeWindowComment(windowCommentID uint, userID uint) error {
	return s.CancelLikeContent(checkWindowCommentExists, consts.ContentType_WindowComment, windowCommentID, userID)
}

// TODO: Reply
