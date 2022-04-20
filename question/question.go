package question

type (
	Question struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}
	Result struct {
		Score float32 `json:"score"`
	}
)
