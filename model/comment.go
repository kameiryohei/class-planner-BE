package model

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	PlanID    uint      `json:"plan_id"`
	UserID    *uint     `json:"user_id"` // 認証ユーザーの場合のみ設定
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user" gorm:"foreignKey:UserID"`
	Plan      Plan      `json:"plan" gorm:"foreignKey:PlanID"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	PlanID    uint      `json:"plan_id"`
	UserID    *uint     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
