package usecase

import (
	"backend/model"
	"backend/repository"
)

type ICommentUsecase interface {
	CreateComment(comment *model.Comment) (model.CommentResponse, error)
	GetCommentsByPlanID(planID uint) ([]model.CommentResponse, error)
	GetCommentsByUserID(userID uint) ([]model.CommentResponse, error)
	DeleteComment(commentID uint, userID *uint) error
}

type commentUsecase struct {
	cr repository.ICommentRepository
}

func NewCommentUsecase(cr repository.ICommentRepository) ICommentUsecase {
	return &commentUsecase{cr: cr}
}

func (cu *commentUsecase) CreateComment(comment *model.Comment) (model.CommentResponse, error) {
	if err := cu.cr.CreateComment(comment); err != nil {
		return model.CommentResponse{}, err
	}

	response := model.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		PlanID:    comment.PlanID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	return response, nil
}

func (cu *commentUsecase) GetCommentsByPlanID(planID uint) ([]model.CommentResponse, error) {
	comments, err := cu.cr.GetCommentsByPlanID(planID)
	if err != nil {
		return nil, err
	}

	responses := make([]model.CommentResponse, len(comments))
	for i, comment := range comments {
		responses[i] = model.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			PlanID:    comment.PlanID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
		}
	}

	return responses, nil
}

func (cu *commentUsecase) GetCommentsByUserID(userID uint) ([]model.CommentResponse, error) {
	comments, err := cu.cr.GetCommentsByUserID(userID)
	if err != nil {
		return nil, err
	}

	responses := make([]model.CommentResponse, len(comments))
	for i, comment := range comments {
		responses[i] = model.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			PlanID:    comment.PlanID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
		}
	}

	return responses, nil
}

func (cu *commentUsecase) DeleteComment(commentID uint, userID *uint) error {
	return cu.cr.DeleteComment(commentID, userID)
}
