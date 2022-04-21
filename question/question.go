package question

type (
	Question struct {
		ID      string `json:"id"`
		Text    string `json:"text"`
		Correct string `json:"answer"`
	}
	Answer struct {
		QuestionID string `json:"question_id"`
		Text       string `json:"text"`
	}
	AnswerResult struct {
		QuestionID string  `json:"question_id"`
		Score      float32 `json:"score"`
	}
)

func NewAnswerResult(questionID string, score float32) *AnswerResult {
	return &AnswerResult{QuestionID: questionID, Score: score}
}
