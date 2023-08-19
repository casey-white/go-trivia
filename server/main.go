package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-trivia/models/quiz_response"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB = initDB()

func main() {
	defer db.Close()

	http.HandleFunc("/quiz", handleQuiz)

	data, err := quiz_response.LoadOneQuizResponse("https://opentdb.com/api.php?amount=5")
	if err != nil {
		log.Fatal(err)
	}

	// Print the data
	for _, quiz := range data.Results {
		quiz_response.SaveQuiz(db, quiz)
	}

	// quiz := quiz_response.LoadOneQuiz(db, "multiple")

	// fmt.Println(quiz)
	http.ListenAndServe(":8080", nil)
}

func handleQuiz(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getQuiz(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getQuiz(w http.ResponseWriter, r *http.Request) {
	data, err := quiz_response.LoadOneQuiz(db, "multiple")

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

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
