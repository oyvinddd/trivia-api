package tapi

import (
	"context"
	"encoding/json"
	"github.com/oyvinddd/trivia-api/question"
	"io"
)

type TriviaAPI struct {
	service question.Service
}

func New(ctx context.Context) *TriviaAPI {
	return &TriviaAPI{service: question.NewService(ctx)}
}

func (tapi TriviaAPI) GetDailyQuestion(ctx context.Context) (*question.Question, error) {
	return tapi.service.GetDailyQuestion(ctx)
}

func (tapi TriviaAPI) SubmitAnswer(ctx context.Context, body io.ReadCloser) (*question.AnswerResult, error) {
	var answer question.Answer
	if err := json.NewDecoder(body).Decode(&answer); err != nil {
		return nil, err
	}
	return tapi.service.SubmitAnswer(ctx, answer)
}
