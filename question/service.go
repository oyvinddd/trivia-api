package question

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

const credentialsFilePath string = "trivia-app-347815-4e057c694d56.json"

type (
	Service interface {
		GetDailyQuestion(ctx context.Context) (*Question, error)

		GetQuestionByID(ctx context.Context, id string) (*Question, error)

		SubmitAnswer(ctx context.Context, answer Answer) (*AnswerResult, error)
	}

	firebaseService struct {
		app *firebase.App
	}
)

func NewService(ctx context.Context) Service {
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
	return nil, nil
}
