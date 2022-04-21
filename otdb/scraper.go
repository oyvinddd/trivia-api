package otdb

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

// instantiates a new Open Trivia DB question
func newQuestion(category, difficulty, text, correct string, incorrect []string) otdbQuestion {
	return otdbQuestion{Category: category, Difficulty: difficulty, Question: text, CorrectAnswer: correct, IncorrectAnswers: incorrect}
}

// Run fetches questions from the Open Trivia API and writes them to a CSV file
func (scraper *Scraper) Run(requests int) error {
	//endTime := time.Now().Add(time.Second * 60) // run for one minute
	apiURL := createOpenTriviaURL(50) // 50 is the max amount the API will give us
	for counter := 0; counter < requests; counter++ {
		res, err := http.Get(apiURL)
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
func (scraper *Scraper) LoadFromFile(filename string) error {
	filePath := fmt.Sprintf("%s/%s", csvDirectory, filename)
	path, err := filepath.Abs(filePath)
	fh, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	lineNumber := 0
	for scanner.Scan() {
		// ignore CSV file header
		if lineNumber == 0 {
			continue
		}
		lineNumber++
		line := scanner.Text()
		question, err := questionFromString(line, "|")
		if err != nil {
			continue
		}
		scraper.AddQuestion(question)
	}
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

// AddQuestion adds a question to the list of questions. dupes are ignored.
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
	return fmt.Sprintf("%s?amount=%d&type=multiple", otdbBaseURL, noOfQuestions)
}

func questionFromString(line, separator string) (otdbQuestion, error) {
	parts := strings.Split(line, separator)
	if len(parts) != 7 {
		return otdbQuestion{}, errors.New("invalid question")
	}
	category := parts[0]
	difficulty := parts[1]
	text := parts[2]
	correct := parts[4]
	incorrect := []string{parts[5], parts[6], parts[7]}
	question := newQuestion(category, difficulty, text, correct, incorrect)
	return question, nil
}
