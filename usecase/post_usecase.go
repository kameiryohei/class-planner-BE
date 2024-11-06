package usecase

import (
	"backend/model"
	"backend/repository"
)

type IPostUsecase interface {
	GetAllPosts(author_id uint) ([]model.PostResponse, error)
	GetPostByID(planId uint) ([]model.PostResponse, error)
	CreatePost(post *model.Post) (model.PostResponse, error)
	DeletePostByID(postId uint) error
}

type postUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(pr repository.IPostRepository) IPostUsecase {
	return &postUsecase{pr}
}

func (pu *postUsecase) GetAllPosts(author_id uint) ([]model.PostResponse, error) {
	posts := []model.Post{}
	if err := pu.pr.GetAllPosts(&posts, author_id); err != nil {
		return nil, err
	}
	resPosts := []model.PostResponse{}
	for _, v := range posts {
		p := model.PostResponse{
			ID:        v.ID,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		}
		resPosts = append(resPosts, p)
	}
	return resPosts, nil

}

func (pu *postUsecase) GetPostByID(planId uint) ([]model.PostResponse, error) {
	posts := []model.Post{}
	if err := pu.pr.GetPostByID(&posts, planId); err != nil {
		return nil, err
	}
	resPosts := []model.PostResponse{}
	for _, v := range posts {
		p := model.PostResponse{
			ID:        v.ID,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		}
		resPosts = append(resPosts, p)
	}

	return resPosts, nil
}

func (pu *postUsecase) CreatePost(post *model.Post) (model.PostResponse, error) {
	if err := pu.pr.CreatePost(post); err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:        post.ID,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) DeletePostByID(postId uint) error {
	if err := pu.pr.DeletePostByID(postId); err != nil {
		return err
	}
	return nil
}
