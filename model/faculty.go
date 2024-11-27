package model

type Faculty struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type FacultyResponse struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
