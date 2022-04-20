package question

type Service interface {
	GetDailyQuestion() (*Question, error)

	GetQuestion(id string) (*Question, error)

	SubmitAnswer()
}
