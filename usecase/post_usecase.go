package usecase

import (
	"backend/model"
	"backend/repository"
)

type IPostUsecase interface {
	GetAllPosts(userId uint) ([]model.PostResponse, error)
	GetPostByID(postId uint, id uint) (model.PostResponse, error)
	CreatePost(post model.Post) (model.PostResponse, error)
	DeletePostByID(postId uint) error
}

type postUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(pr repository.IPostRepository) IPostUsecase {
	return &postUsecase{pr}
}

func (pu *postUsecase) GetAllPosts(userId uint) ([]model.PostResponse, error) {
	posts := []model.Post{}
	if err := pu.pr.GetAllPosts(&posts, userId); err != nil {
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

func (pu *postUsecase) GetPostByID(id uint, postId uint) (model.PostResponse, error) {
	post := model.Post{}
	if err := pu.pr.GetPostByID(&post, id, postId); err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:        post.ID,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) CreatePost(post model.Post) (model.PostResponse, error) {
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
