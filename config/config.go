package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RedisHost  string
	RedisPort  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func (c *Config) migrateUp() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=require",
		c.DBUser, c.DBPassword, c.DBName, c.DBHost, c.DBPort)

	os.Setenv("GOOSE_DRIVER", "postgres")
	os.Setenv("GOOSE_DBSTRING", dsn)

	migrationsDir := "./db"

	cmd := exec.Command("goose", "-dir", migrationsDir, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Running database migrations...")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully.")
}

func (c *Config) ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	c.migrateUp()

	return db
}
