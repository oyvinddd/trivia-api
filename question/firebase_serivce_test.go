package question

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oyvinddd/trivia-api/config"
	"log"
	"os"
	"testing"
)

func createAndInitFirebaseService() Service {
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatalln(err)
	}
	cfg := config.Firebase(
		os.Getenv("FB_TYPE"), os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"), os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"), os.Getenv("FB_CLIENT_ID"))
	return NewService(context.TODO(), cfg)
}

func TestFirebaseService_GetDailyQuestions(t *testing.T) {
	service := createAndInitFirebaseService()
	container, err := service.GetDailyQuestions(context.TODO())
	if err != nil {
		t.Error(err)
	}
	if len(container.Questions) != 5 {
		t.Errorf("# of questions on list is not correct: %d\n", len(container.Questions))
	}
}

func TestFirebaseService_SubmitAnswer(t *testing.T) {
	service := createAndInitFirebaseService()

	a1 := Answer{QuestionID: 1, Text: "Donald Trump"}
	a2 := Answer{QuestionID: 1, Text: "donald trump"}

	res, err := service.SubmitAnswer(context.TODO(), a1)
	if err != nil {
		t.Error(err)
	}
	res, err = service.SubmitAnswer(context.TODO(), a2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
