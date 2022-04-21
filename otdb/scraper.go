package otdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const (
	otdbBaseURL  string = "https://opentdb.com/api.php"
	csvDirectory string = "data"
)

type (
	// Scraper wraps our API scraping functionality
	Scraper struct {
		Questions []otdbQuestion
	}

	// otdbQuestion represents the raw Open Trivia DB question
	otdbQuestion struct {
		Category         string   `json:"category"`
		QuestionType     string   `json:"type"`
		Difficulty       string   `json:"difficulty"`
		Question         string   `json:"question"`
		CorrectAnswer    string   `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
	}

	// OTDBResponse response body we get from Open Trivia DB
	otdbResponse struct {
		Code      int            `json:"response_code"`
		Questions []otdbQuestion `json:"results"`
	}
)

// New creates a new Scraper instance
func New() *Scraper {
	return &Scraper{make([]otdbQuestion, 0)}
}

// Run fetches questions from the Open Trivia API and writes them to a CSV file
func (scraper *Scraper) Run(requests int) error {
	//endTime := time.Now().Add(time.Second * 60) // run for one minute
	for counter := 0; counter < requests; counter++ {
		res, err := http.Get(createOpenTriviaURL(50)) // 50 is the max amount the API will give us
		if err != nil || res.StatusCode != http.StatusOK {
			return err
		}
		var otdb otdbResponse
		if err := json.NewDecoder(res.Body).Decode(&otdb); err != nil {
			return err
		}
		for _, question := range otdb.Questions {
			scraper.AddQuestion(question)
		}
	}
	return nil
}

// LoadFromFile loads questions from a given CSV file
func (scraper Scraper) LoadFromFile(filename string) error {
	return nil
}

// WriteToFile writes (appends) questions to a given CSV file
func (scraper Scraper) WriteToFile(filename string) error {
	if scraper.HasNoQuestions() {
		return nil
	}
	filePath := fmt.Sprintf("%s/%s", csvDirectory, filename)
	path, err := filepath.Abs(filePath)
	fh, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer fh.Close()
	for _, question := range scraper.Questions {
		line := fmt.Sprintf("%s\n", question.ToCSVFormat("|"))
		_, err := fh.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddQuestion adds a question to the list of questions
func (scraper *Scraper) AddQuestion(question otdbQuestion) {
	scraper.Questions = append(scraper.Questions, question)
}

func (scraper Scraper) NoOfQuestions() int {
	return len(scraper.Questions)
}

// HasNoQuestions checks if scraper has any questions
func (scraper Scraper) HasNoQuestions() bool {
	return len(scraper.Questions) == 0
}

// PrintQuestions prints questions to standard output
func (scraper Scraper) PrintQuestions() {
	for _, question := range scraper.Questions {
		fmt.Println(question.ToCSVFormat("|"))
	}
}

// ToCSVFormat use this function when writing the question to a CSV file
func (question otdbQuestion) ToCSVFormat(separator string) string {
	var incorrectStr string
	for _, incorrect := range question.IncorrectAnswers {
		incorrectStr += fmt.Sprintf("%s%s", incorrect, separator)
	}
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
		question.Category,
		separator,
		question.Difficulty,
		separator,
		question.Question,
		separator,
		question.CorrectAnswer,
		separator,
		incorrectStr)
}

func createOpenTriviaURL(noOfQuestions int) string {
	return fmt.Sprintf("%s?amount=%d&encode=url3986", otdbBaseURL, noOfQuestions)
}
