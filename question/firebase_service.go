package question

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/oyvinddd/trivia-api/config"
	"github.com/oyvinddd/trivia-api/levenshtein"
	"google.golang.org/api/option"
	"log"
	"strings"
)

// this struct implements our main Service interface
type firebaseService struct {
	app *firebase.App
}

func NewService(ctx context.Context, cfg config.Config) Service {
	credentials := option.WithCredentialsJSON(cfg.Bytes())
	app, err := firebase.NewApp(ctx, nil, credentials)
	if err != nil {
		log.Fatalln(err)
	}
	return &firebaseService{app: app}
}

func (service firebaseService) GetDailyQuestion(ctx context.Context) (*Question, error) {
	//client, err := service.app.Firestore(ctx)
	//defer client.Close()
	// TODO: create logic for daily question
	return service.GetQuestionByID(ctx, "1")
}

func (service firebaseService) GetQuestionByID(ctx context.Context, id string) (*Question, error) {
	client, err := service.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	snapshot, err := client.Collection("questions").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var question Question
	if err := snapshot.DataTo(&question); err != nil {
		return nil, err
	}
	return &question, nil
}

func (service firebaseService) SubmitAndEvaluateAnswer(ctx context.Context, answer Answer) (*AnswerResult, error) {
	question, err := service.GetQuestionByID(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}
	score := service.EvaluateAnswer(*question, answer)
	return NewAnswerResult(question.ID, score), nil
}

func (service firebaseService) EvaluateAnswer(question Question, answer Answer) float32 {
	answerLower := strings.ToLower(answer.Text)
	correctLower := strings.ToLower(question.Answer)
	// first check if answer needs to be matched exactly
	if question.NeedsExactMatch() {
		if answerLower == correctLower {
			return 100.0
		}
		return 0.0
	}
	// if we don't require an exact match, use the Edit Distance algorithm
	// to calculate a score for the user
	return levenshtein.Calculate(answerLower, correctLower)
}
