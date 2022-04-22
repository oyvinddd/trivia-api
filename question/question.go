package question

import (
	"fmt"
	"strings"
)

type (
	// Question represents our domain-specific question model
	Question struct {
		ID       string `json:"id"`
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	// Answer is represents a user's given answer to a question
	Answer struct {
		QuestionID string `json:"question_id"`
		Text       string `json:"text"`
	}

	// AnswerResult represents the quality (0.0 to 100.0) of a particular answer to a question
	AnswerResult struct {
		QuestionID string  `json:"question_id"`
		Score      float32 `json:"score"`
	}
)

// New creates a new question
func New(id, text, answer string) *Question {
	return &Question{ID: id, Question: text, Answer: answer}
}

// NewAnswerResult creates a new answer result
func NewAnswerResult(questionID string, score float32) *AnswerResult {
	return &AnswerResult{QuestionID: questionID, Score: score}
}

// NeedsExactMatch answers that are only one word or are less than or equal to 6
// characters in length requires an exact match in order for the user to be correct
func (question Question) NeedsExactMatch() bool {
	return len(question.Answer) <= 6 || len(strings.Fields(question.Answer)) == 1
}

// String returns a string representation of the given question
func (question Question) String() string {
	return fmt.Sprintf("[%s] Question: %s, answer: %s", question.ID, question.Question, question.Answer)
}
