package model

type Faculty struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"not null"`
	UniversityID uint   `json:"university_id" gorm:"not null"`

	University  University   `json:"university" gorm:"foreignKey:UniversityID"`
	Users       []User       `json:"users" gorm:"foreignKey:FacultyID"`
	Departments []Department `json:"departments" gorm:"foreignKey:FacultyID"`
}
