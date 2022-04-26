package question

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/oyvinddd/trivia-api/config"
	"github.com/oyvinddd/trivia-api/levenshtein"
	"google.golang.org/api/option"
	"log"
	"math/rand"
	"strings"
)

const noOfQuestions int = 874 // TODO: import more questions

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
	client, err := service.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	randomQuestionID := randomNumber(1, noOfQuestions)
	return service.GetQuestionByID(ctx, randomQuestionID)
}

func (service firebaseService) GetDailyQuestions(context.Context) ([]Question, error) {
	log.Fatalln("GetDailyQuestions has not been implemented yet")
	return nil, nil
}

func (service firebaseService) GetQuestionByID(ctx context.Context, id int) (*Question, error) {
	client, err := service.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	questions := client.Collection("questions")
	iter := questions.Where("ID", "==", id).Limit(1).Documents(ctx)
	// we don't need to iterate here since we're only interested in the first object
	snapshot, err := iter.Next()
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

// returns a random number between min and max
func randomNumber(min, max int) int {
	return rand.Intn((max - min) + min)
}
