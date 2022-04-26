package tapi

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oyvinddd/trivia-api/config"
	"log"
	"os"
	"testing"
)

func createAndInitTriviaAPI() *TriviaAPI {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalln(err)
	}
	cfg := config.Firebase(
		os.Getenv("FB_TYPE"), os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"), os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"), os.Getenv("FB_CLIENT_ID"))
	return New(context.TODO(), cfg)
}

func TestGetDailyQuestions(t *testing.T) {
	questions, err := createAndInitTriviaAPI().GetDailyQuestions()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(questions)
}

func TestGetQuestionByID(t *testing.T) {
	question, err := createAndInitTriviaAPI().GetQuestionByID(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(question.String())
}
