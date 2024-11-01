package model

type Post struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Content   *string `json:"content"`
	PlanID    *uint   `json:"plan_id"`
	CreatedAt *string `json:"created_at"`

	Author   User  `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID uint  `json:"author_id"`
	Plan     *Plan `json:"plan" gorm:"foreignKey:PlanID"`
}
type PostResponse struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Content   *string `json:"content"`
	CreatedAt *string `json:"created_at"`
}
