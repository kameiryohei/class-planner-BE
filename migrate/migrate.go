package main

import (
	"backend/db"
	"backend/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{},
		&model.University{},
		&model.Course{},
		&model.Department{},
		&model.Faculty{},
		&model.FavoritePlan{},
		&model.Plan{},
		&model.Post{},
		&model.Comment{},
		&model.TimetableAnalysis{},
		&model.ExtractedClass{},
	)
}
