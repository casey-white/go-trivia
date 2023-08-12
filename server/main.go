package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"go-trivia/models/quiz_response"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	db := initDB()
	defer db.Close()

	data, err := quiz_response.LoadOneQuizResponse("https://opentdb.com/api.php?amount=1")
	if err != nil {
		log.Fatal(err)
	}

	// Print the data
	for _, quiz := range data.Results {
		quiz_response.SaveQuiz(db, quiz)
	}

	quiz := quiz_response.LoadOneQuiz(db, "multiple")

	fmt.Println(quiz)
}

func initDB() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=db port=5432 sslmode=disable", dbUser, dbName, dbPassword)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
