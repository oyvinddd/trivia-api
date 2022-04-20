package question

type (
	Question struct {
		ID   string `json:"id"`
		Text string `json:"text"`
		Correct string `json:"correct_answer"`
	}
	Result struct {
		Score float32 `json:"score"`
	}
)
