package otdb

import (
	"testing"
)

func TestFetchQuestions(t *testing.T) {
	scraper := New()
	scraper.Run(10)
	//scraper.PrintQuestions()
	err := scraper.WriteToFile("questions.csv")
	if err != nil {
		t.Error(err)
	}
}
