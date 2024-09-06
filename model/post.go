package model

type Post struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Content   *string `json:"content"`
	Published bool    `json:"published" gorm:"default:false"`
	AuthorID  *uint   `json:"author_id"`
	PlanID    *uint   `json:"plan_id"`

	Author User  `json:"author" gorm:"foreignKey:AuthorID"`
	Plan   *Plan `json:"plan" gorm:"foreignKey:PlanID"`
}
