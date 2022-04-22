package otdb

import (
	"testing"
)

func TestFetchQuestions(t *testing.T) {
	scraper := New()
	scraper.Run(20)
	err := scraper.WriteToFile("questions.csv")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadQuestionsFromFile(t *testing.T) {
	scraper := New()
	scraper.LoadFromFile("questions.csv")
	scraper.PrintQuestions()
}
