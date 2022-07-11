package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := "mysql"
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, _ = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database", Dbdriver)
	}

	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&User{})
}
