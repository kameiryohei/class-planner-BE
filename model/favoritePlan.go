package model

type FavoritePlan struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`
	PlanID uint `json:"plan_id" gorm:"not null"`

	User User `json:"user" gorm:"foreignKey:UserID"`
	Plan Plan `json:"plan" gorm:"foreignKey:PlanID"`
}

type FavoritePlanResponse struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	PlanID uint `json:"plan_id"`
}
