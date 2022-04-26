package question

import (
	"context"
)

type Service interface {

	// GetDailyQuestion gets a daily question from the service
	GetDailyQuestion(ctx context.Context) (*Question, error)

	// GetDailyQuestions gets a set of daily questions from the service
	GetDailyQuestions(ctx context.Context) ([]Question, error)

	// GetQuestionByID gets a question by a given ID from the service
	GetQuestionByID(ctx context.Context, id int) (*Question, error)

	// SubmitAndEvaluateAnswer submits an answer for a given question to the service
	SubmitAndEvaluateAnswer(ctx context.Context, answer Answer) (*AnswerResult, error)

	// EvaluateAnswer evaluates the correctness of the given answer
	EvaluateAnswer(question Question, answer Answer) float32
}
