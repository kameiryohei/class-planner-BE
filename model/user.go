package model

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	UniversityID *uint  `json:"university_id"`
	FacultyID    *uint  `json:"faculty_id"`
	DepartmentID *uint  `json:"department_id"`
	Grade        *uint  `json:"grade"`

	University *University    `json:"university" gorm:"foreignKey:UniversityID"`
	Faculty    *Faculty       `json:"faculty" gorm:"foreignKey:FacultyID"`
	Department *Department    `json:"department" gorm:"foreignKey:DepartmentID"`
	Posts      []Post         `json:"posts" gorm:"foreignKey:AuthorID"`
	Plans      []Plan         `json:"plans" gorm:"foreignKey:UserID"`
	Favorites  []FavoritePlan `json:"favorites" gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	Email      string      `json:"email" gorm:"unique"`
	University *University `json:"university"`
	Faculty    *Faculty    `json:"faculty"`
	Department *Department `json:"department"`
}

type GoogleUserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
