package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *model.Comment) error
	GetCommentsByPlanID(planID uint) ([]model.Comment, error)
	GetCommentsByUserID(userID uint) ([]model.Comment, error)
	DeleteComment(commentID uint, userID *uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db: db}
}

func (cr *commentRepository) CreateComment(comment *model.Comment) error {
	return cr.db.Create(comment).Error
}

func (cr *commentRepository) GetCommentsByPlanID(planID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := cr.db.Preload("User").Where("plan_id = ?", planID).Order("created_at desc").Find(&comments).Error
	return comments, err
}

func (cr *commentRepository) GetCommentsByUserID(userID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := cr.db.Preload("Plan").Where("user_id = ?", userID).Order("created_at desc").Find(&comments).Error
	return comments, err
}

func (cr *commentRepository) DeleteComment(commentID uint, userID *uint) error {
	query := cr.db.Where("id = ?", commentID)
	if userID != nil {
		query = query.Where("user_id = ?", userID)
	}
	result := query.Delete(&model.Comment{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
