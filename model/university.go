package model

type University struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`

	Faculties []Faculty `json:"faculties" gorm:"foreignKey:UniversityID"`
}
