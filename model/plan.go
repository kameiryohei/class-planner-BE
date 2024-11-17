package model

import "time"

type Plan struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Content   *string   `json:"content"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User      User           `json:"user" gorm:"foreignKey:UserID"`
	Courses   []Course       `json:"courses" gorm:"foreignKey:PlanID"`
	Posts     []Post         `json:"posts" gorm:"foreignKey:PlanID"`
	Favorites []FavoritePlan `json:"favorites" gorm:"foreignKey:PlanID"`
}

type PlanResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   *string   `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type PlanDetailResponse struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	Title     string           `json:"title"`
	Content   *string          `json:"content"`
	UserID    uint             `json:"user_id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	User      UserResponse     `json:"user"`
	Courses   []CourseResponse `json:"courses" gorm:"foreignKey:PlanID"`
	Posts     []Post           `json:"posts" gorm:"foreignKey:PlanID"`
	Favorites []FavoritePlan   `json:"favorites" gorm:"foreignKey:PlanID"`
}
