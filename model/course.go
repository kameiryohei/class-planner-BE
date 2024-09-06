package model

import "time"

type Course struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"not null"`
	Content   *string    `json:"content"`
	PlanID    uint       `json:"plan_id" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`

	Plan Plan `json:"plan" gorm:"foreignKey:PlanID"`
}
