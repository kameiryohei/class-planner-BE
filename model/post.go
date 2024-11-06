package model

import "time"

type Post struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Content   *string    `json:"content"`
	PlanID    *uint      `json:"plan_id"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	Author    User       `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID  uint       `json:"author_id"`
	Plan      *Plan      `json:"plan" gorm:"foreignKey:PlanID"`
}
type PostResponse struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Content   *string    `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
}
