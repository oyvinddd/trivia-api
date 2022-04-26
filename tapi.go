package tapi

import (
	"context"
	"encoding/json"
	"github.com/oyvinddd/trivia-api/config"
	"github.com/oyvinddd/trivia-api/question"
	"io"
)

type TriviaAPI struct {
	ctx     context.Context
	service question.Service
}

// New creates a new TriviaAPI instance
func New(ctx context.Context, cfg config.Config) *TriviaAPI {
	return &TriviaAPI{ctx: ctx, service: question.NewService(ctx, cfg)}
}

// GetDailyQuestion gets the daily question from the API
func (tapi TriviaAPI) GetDailyQuestion() (*question.Question, error) {
	return tapi.service.GetDailyQuestion(tapi.ctx)
}

// GetQuestionByID gets a question with a specific ID from the API
func (tapi TriviaAPI) GetQuestionByID(id int) (*question.Question, error) {
	return tapi.service.GetQuestionByID(tapi.ctx, id)
}

// SubmitAnswer submits an answer for a given question to the API
func (tapi TriviaAPI) SubmitAnswer(body io.ReadCloser) (*question.AnswerResult, error) {
	var answer question.Answer
	if err := json.NewDecoder(body).Decode(&answer); err != nil {
		return nil, err
	}
	return tapi.service.SubmitAndEvaluateAnswer(tapi.ctx, answer)
}
