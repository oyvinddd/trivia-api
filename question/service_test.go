package question

import (
	"context"
	"fmt"
	"github.com/oyvinddd/trivia-api/config"
	"testing"
)

func TestGetQuestionByID(t *testing.T) {
	ctx := context.TODO()

	cfg, err := config.FromEnvFile("../local.env")
	if err != nil {
		t.Error(err)
	}
	service := NewService(ctx, cfg)

	question, err := service.GetQuestionByID(ctx, "1")

	if err != nil {
		t.Error(err)
	}
	if question == nil {
		t.Error("returned question was nil")
	}
	fmt.Println(question.String())
}
