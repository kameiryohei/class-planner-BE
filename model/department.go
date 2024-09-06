package model

type Department struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	FacultyID uint   `json:"faculty_id" gorm:"not null"`

	Faculty Faculty `json:"faculty" gorm:"foreignKey:FacultyID"`
	Users   []User  `json:"users" gorm:"foreignKey:DepartmentID"`
}
