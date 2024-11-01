package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
)

type IPostRepository interface {
	GetAllPosts(posts *[]model.Post, userId uint) error //ユーザーが作成した全ての投稿を取得
	GetPostByID(post *model.Post, id uint, postId uint) error
	CreatePost(post model.Post) error
	DeletePostByID(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{db}
}

// すべての投稿を取得
func (pr *postRepository) GetAllPosts(posts *[]model.Post, userId uint) error {
	if err := pr.db.Joins("User").Where("author_id = ?", userId).Order("created_at").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

// 投稿IDで投稿を取得
func (pr *postRepository) GetPostByID(post *model.Post, courseId uint, postId uint) error {
	if err := pr.db.Joins("Plan").Where("id = ?", courseId).First(post, postId).Error; err != nil {
		return err
	}
	return nil
}

// 投稿を作成
func (pr *postRepository) CreatePost(post model.Post) error {
	if err := pr.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

// 投稿を削除
func (pr *postRepository) DeletePostByID(postId uint) error {
	result := pr.db.Where("id = ?", postId).Delete(&model.Post{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
