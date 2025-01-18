package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

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
func (c *Config) GenerateMocks() {
	// Use 'go list' with a format string to get all Go package directories in the project
	cmd := exec.Command("go", "list", "-f", "{{.Dir}}", "./...")
	cmd.Dir = "." // Set the working directory to your project root

	// Capture the output from the `go list` command
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to list Go files: %v", err)
	}

	// Split the output into directories (each directory corresponds to a Go package)
	dirs := strings.Split(string(output), "\n")

	// Create a mock folder if it doesn't exist
	mockFolder := "./mock"
	err = os.MkdirAll(mockFolder, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create mock folder: %v", err)
	}

	// Loop through each directory and find .go files in it
	for _, dir := range dirs {
		if dir == "" {
			continue
		}

		// Look for all Go files within this directory
		files, err := os.ReadDir(dir)
		if err != nil {
			log.Fatalf("Failed to read directory %s: %v", dir, err)
		}

		// Loop through the files in the directory
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".go") {
				goFilePath := fmt.Sprintf("%s/%s", dir, file.Name())

				// Check if the file contains any interfaces by reading it
				if containsInterface(goFilePath) {
					// Use the mock folder to store the mock files
					mockFileName := fmt.Sprintf("%s/%s_mock.go", mockFolder, strings.TrimSuffix(file.Name(), ".go"))

					// Generate mocks using mockgen
					cmd := exec.Command("mockgen", "-source="+goFilePath, "-destination="+mockFileName, "-package=mock")
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					err := cmd.Run()
					if err != nil {
						log.Fatalf("Failed to generate mock for %s: %v", goFilePath, err)
					}
					log.Printf("Mock generated for %s", goFilePath)
				}
			}
		}
	}
}

// containsInterface checks if the Go file contains at least one interface definition
func containsInterface(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	// Check for interface keyword in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "interface") {
			return true
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Fatalf("Error reading file %s: %v", filePath, err)
	}

	return false
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
