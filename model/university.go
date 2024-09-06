package model

type University struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`

	Users     []User    `json:"users" gorm:"foreignKey:UniversityID"`
	Faculties []Faculty `json:"faculties" gorm:"foreignKey:UniversityID"`
}
