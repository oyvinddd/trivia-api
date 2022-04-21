package question

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/oyvinddd/trivia-api/levenshtein"
	"google.golang.org/api/option"
	"log"
	"strings"
)

const credentialsFilePath string = "trivia-app-347815-4e057c694d56.json"

type (
	Service interface {
		GetDailyQuestion(ctx context.Context) (*Question, error)

		GetQuestionByID(ctx context.Context, id string) (*Question, error)

		SubmitAnswer(ctx context.Context, answer Answer) (*AnswerResult, error)

		EvaluateAnswer(question Question, answer Answer) float32
	}

	firebaseService struct {
		app *firebase.App
	}
)

func NewService(ctx context.Context, cfg Config) Service {
	// TODO: initialize Firebase app from config instead
	sa := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	return &firebaseService{app: app}
}

func (service firebaseService) GetDailyQuestion(ctx context.Context) (*Question, error) {
	client, err := service.app.Firestore(ctx)
	defer client.Close()
	// TODO: fetch daily question
	return nil, err
}

func (service firebaseService) GetQuestionByID(ctx context.Context, id string) (*Question, error) {
	client, err := service.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	_ = client.Collection("questions").Doc(id)
	return nil, nil
}

func (service firebaseService) SubmitAnswer(ctx context.Context, answer Answer) (*AnswerResult, error) {
	question, err := service.GetQuestionByID(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}
	score := service.EvaluateAnswer(*question, answer)
	return NewAnswerResult(question.ID, score), nil
}

func (service firebaseService) EvaluateAnswer(question Question, answer Answer) float32 {
	answerLower := strings.ToLower(answer.Text)
	correctLower := strings.ToLower(question.Correct)
	// first check if answer needs to be matched exactly
	if needsExactMatch(correctLower) {
		if answerLower == correctLower {
			return 100.0
		}
		return 0.0
	}
	// if we don't require an exact match, use the Edit Distance algorithm
	// to calculate a score for the user
	return levenshtein.Calculate(answerLower, correctLower)
}

// answers that are only one word or are less than or equal to 6 characters in length
// requires an exact match in order for the user to be correct
func needsExactMatch(answer string) bool {
	return len(answer) <= 6 || len(strings.Fields(answer)) == 1
}
