package question

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/oyvinddd/trivia-api/config"
	"github.com/oyvinddd/trivia-api/levenshtein"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	noOfQuestions   int = 874 // TODO: import more questions
	questionsPerDay int = 5
)

var referenceDate = time.Date(2022, 4, 28, 0, 0, 0, 0, time.UTC)

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

func (service firebaseService) GetDailyQuestions(ctx context.Context) ([]Question, error) {
	client, err := service.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	questionIDs := calculateDailyQuestionIDs()
	questionsRef := client.Collection("questions")
	iter := questionsRef.Where("ID", "in", questionIDs).Documents(ctx)
	dailyQuestions := make([]Question, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var question Question
		if err := doc.DataTo(&question); err != nil {
			return nil, err
		}
		dailyQuestions = append(dailyQuestions, question)
	}
	return dailyQuestions, nil
}

func (service firebaseService) GetQuestionByID(ctx context.Context, id int) (*Question, error) {
	client, err := service.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	questionsRef := client.Collection("questions")
	iter := questionsRef.Where("ID", "==", id).Limit(1).Documents(ctx)
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

func (service firebaseService) GetRandomQuestion(ctx context.Context) (*Question, error) {
	return service.GetQuestionByID(ctx, randomNumber(1, noOfQuestions))
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
	// do some preliminary operations on the answers before we compare them
	answerLower := strings.ToLower(strings.TrimSpace(answer.Text))
	correctLower := strings.ToLower(strings.TrimSpace(question.Answer))
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

// calculates the 5 question IDs belonging to the current day
func calculateDailyQuestionIDs() []int {
	// the day # since we started the trivia
	currentDay := time.Now().Sub(referenceDate).Hours() / 24
	// skip ahead by 5 each day, since we're serving up 5 questions every day
	// use modulus here to wrap around once the IDs start exceeding # of questions
	id1 := (int(currentDay) * questionsPerDay) % noOfQuestions
	return []int{id1 + 0, id1 + 1, id1 + 2, id1 + 3, id1 + 4}
}

// returns a random number between min and max
func randomNumber(min, max int) int {
	return rand.Intn((max - min) + min)
}
