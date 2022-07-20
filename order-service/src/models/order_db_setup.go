package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "scalable-final-proj/order-service"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	Dbdriver := "mysql"
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	// DBURL := "root:rootpass@tcp(0.0.0.0:3307)/scalabruh_final_db?charset=utf8&parseTime=True&loc=Local"

	DB, _ = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	DB.DB().SetConnMaxLifetime(0)
	DB.AutoMigrate(&Order{})
	// DB.Model(&Product{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}
