package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

type QuizResponse struct {
	ResponseCode int    `json:"response_code"`
	Results      []Quiz `json:"results"`
}

type Quiz struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

func main() {

	connStr := "user=username dbname=dbname password=password host=host port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://opentdb.com/api.php?amount=1", nil)
	if err != nil {
		panic(err)
	}

	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// Check if the response status code is 200
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("status code: %d", resp.StatusCode))
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON data into a map
	var data QuizResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Print the data
	for _, quiz := range data.Results {
		fmt.Println(quiz.Question)
	}
}
