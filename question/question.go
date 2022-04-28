package question

import (
	"fmt"
	"strings"
)

type (
	// Question represents our domain-specific question model
	Question struct {
		ID         int    `json:"id"`
		Category   string `json:"category"`
		Difficulty string `json:"difficulty"`
		Question   string `json:"question"`
		Answer     string `json:"-"`
	}

	// Answer represents a user's given answer to a question
	Answer struct {
		QuestionID int    `json:"question_id"`
		Text       string `json:"text"`
	}

	// AnswerResult represents the quality (0.0 to 100.0) of a particular answer to a question
	AnswerResult struct {
		QuestionID int     `json:"question_id"`
		Score      float32 `json:"score"`
	}
)

// New creates a new question
func New(id int, category, difficulty, text, answer string) *Question {
	return &Question{ID: id, Category: category, Difficulty: difficulty, Question: text, Answer: answer}
}

// NewAnswerResult creates a new answer result
func NewAnswerResult(questionID int, score float32) *AnswerResult {
	return &AnswerResult{QuestionID: questionID, Score: score}
}

// NeedsExactMatch answers that are only one word or are less than or equal to 6
// characters in length requires an exact match in order for the user to be correct
func (question Question) NeedsExactMatch() bool {
	return len(question.Answer) <= 8 || len(strings.Fields(question.Answer)) == 1
}

// String returns a string representation of the given question
func (question Question) String() string {
	return fmt.Sprintf("[%d][%s][%s] Question: %s, answer: %s",
		question.ID,
		question.Category,
		question.Difficulty,
		question.Question,
		question.Answer,
	)
}
