package tapi

import (
	"context"
	"github.com/oyvinddd/trivia-api/question"
)

type TriviaAPI struct {
	service question.Service
}

func New() *TriviaAPI {
	return &TriviaAPI{service: question.NewService()}
}

func (tapi TriviaAPI) GetDailyQuestion(ctx context.Context) (*question.Question, error) {
	return nil, nil
}
