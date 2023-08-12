package models

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
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

func LoadOneQuizResponse(apiURL string) (*QuizResponse, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data QuizResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func SaveQuiz(db *sql.DB, quiz Quiz) error {
	query := `INSERT INTO quizzes (category, type, difficulty, question, correct_answer) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, quiz.Category, quiz.Type, quiz.Difficulty, quiz.Question, quiz.CorrectAnswer)
	return err
}
