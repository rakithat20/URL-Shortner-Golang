package Config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	// Import the pgx driver
	"github.com/joho/godotenv" // For loading .env files
	"github.com/patrickmn/go-cache"
)

var DB *sql.DB
var CACHE = cache.New(6*time.Hour, 12*time.Hour)

// ConnectDB initializes the database connection
func ConnectDB(l *log.Logger) {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		l.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve database credentials from .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build the connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// Open a connection to the database
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		l.Fatalf("Failed to connect to database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		l.Fatalf("Failed to ping database: %v", err)
	}

	l.Println("Database connection established")
}
