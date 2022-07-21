package p_models

import (
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {


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
	// DBURL := "root:rootpass@tcp(0.0.0.0:3308)/scalabruh_final_proj_prod_db?charset=utf8&parseTime=True&loc=Local"

	DB, _ = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&Product{})
}
