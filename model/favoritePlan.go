package model

import "time"

type FavoritePlan struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	PlanID    uint      `json:"plan_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`

	User User `json:"user" gorm:"foreignKey:UserID"`
	Plan Plan `json:"plan" gorm:"foreignKey:PlanID"`
}
