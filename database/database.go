package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

func NewDB() (*gorm.DB, error) {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	env := os.Getenv("ENV")
	DB := os.Getenv("DB")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	CONNECT := DBUser + ":" + DBPass + "@" + env + "/" + DB + "?parseTime=true"

	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		return nil, fmt.Errorf("failed to connect db= %w", err)
	}
	return db, nil
}
