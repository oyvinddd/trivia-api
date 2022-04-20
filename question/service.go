package question

import "context"

type (
	Service interface {
		GetDailyQuestion(ctx context.Context) (*Question, error)

		GetQuestionByID(ctx context.Context, id string) (*Question, error)

		SubmitAnswer(ctx context.Context, answer string) error
	}

	firebaseService struct{}
)

func NewService() Service {
	return &firebaseService{}
}

func (service firebaseService) GetDailyQuestion(ctx context.Context) (*Question, error) {
	return nil, nil
}

func (service firebaseService) GetQuestionByID(ctx context.Context, id string) (*Question, error) {
	return nil, nil
}

func (service firebaseService) SubmitAnswer(ctx context.Context, answer string) error {
	return nil
}
