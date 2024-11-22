package model

type Department struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type DepartmentResponse struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
