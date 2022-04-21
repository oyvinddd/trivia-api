package question

import (
	"context"
	"testing"
)

func TestFetchQuestionByID(t *testing.T) {
	ctx := context.TODO()

	service := NewService(ctx)
	question, err := service.GetQuestionByID(ctx, "1")

	if err != nil {
		t.Error(err)
	}
	if question == nil {
		t.Error("returned question was nil")
	}
}
