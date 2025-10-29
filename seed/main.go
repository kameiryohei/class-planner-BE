package main

import (
	"backend/db"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

func main() {
	if os.Getenv("GO_ENV") != "dev" {
		log.Println("GO_ENV is not dev; skip applying seed data")
		return
	}

	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	log.Println("Applying local seed data")
	if err := seedLocal(dbConn); err != nil {
		log.Fatalf("failed to apply local seed data: %v", err)
	}
	log.Println("Finished applying local seed data")
}

func seedLocal(dbConn *gorm.DB) error {
	seedSQL, err := os.ReadFile("db/seed.sql")
	if err != nil {
		return fmt.Errorf("read seed file: %w", err)
	}

	statements := parseStatements(string(seedSQL))
	if len(statements) == 0 {
		return nil
	}

	return dbConn.Transaction(func(tx *gorm.DB) error {
		for _, stmt := range statements {
			if err := tx.Exec(stmt).Error; err != nil {
				return fmt.Errorf("exec seed statement: %w", err)
			}
		}
		return nil
	})
}

func parseStatements(sql string) []string {
	rawStatements := strings.Split(sql, ";")
	statements := make([]string, 0, len(rawStatements))

	for _, raw := range rawStatements {
		lines := strings.Split(raw, "\n")
		var builder []string

		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, "--") {
				continue
			}
			builder = append(builder, trimmed)
		}

		stmt := strings.Join(builder, " ")
		stmt = strings.TrimSpace(stmt)
		if stmt != "" {
			statements = append(statements, stmt)
		}
	}

	return statements
}
