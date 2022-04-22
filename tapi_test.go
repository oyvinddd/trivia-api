package tapi

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oyvinddd/trivia-api/config"
	"os"
	"testing"
)

func TestGetDailyQuestion(t *testing.T) {
	godotenv.Load("local.env")
	cfg := config.Firebase(
		os.Getenv("FB_TYPE"),
		os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"),
		os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"),
		os.Getenv("FB_CLIENT_ID"))
	question, err := New(context.TODO(), cfg).GetDailyQuestion()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(question)
}
