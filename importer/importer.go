package importer

import (
	"bufio"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/oyvinddd/trivia-api/config"
	"github.com/oyvinddd/trivia-api/question"
	"google.golang.org/api/option"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FirebaseImporter struct {
	ctx context.Context
	app *firebase.App
}

func NewFirebaseImporter(ctx context.Context, cfg config.Config) *FirebaseImporter {
	credentials := option.WithCredentialsJSON(cfg.Bytes())
	app, err := firebase.NewApp(ctx, nil, credentials)
	if err != nil {
		log.Fatalln(err)
	}
	return &FirebaseImporter{ctx: ctx, app: app}
}

func (importer FirebaseImporter) ImportAvailableQuestions() error {
	questions, err := importer.loadQuestionsFromFile("../otdb/data/questions.csv")
	if err != nil {
		return err
	}
	if len(questions) == 0 {
		return errors.New("no questions to import")
	}
	client, err := importer.app.Firestore(importer.ctx)
	if err != nil {
		return err
	}
	counter := 0
	for _, question := range questions {
		_, err := client.Collection("questions").NewDoc().Set(importer.ctx, question)
		//_, err := client.Collection("questions").Doc(question.ID).Set(importer.ctx, question)
		if err != nil {
			log.Fatalf("Failed adding question: %v", err)
		}
		counter++
	}
	fmt.Printf("Successfully imported %d questions to Firestore\n", counter)
	return nil
}

func (importer FirebaseImporter) loadQuestionsFromFile(filePath string) ([]question.Question, error) {
	path, err := filepath.Abs(filePath)
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	var questions = make([]question.Question, 0)
	scanner := bufio.NewScanner(fh)
	isHeader := true
	for scanner.Scan() {
		if isHeader {
			isHeader = false
			continue
		}
		line := scanner.Text()
		question, err := questionFromString(line, "|")
		if err != nil {
			//log.Errorf("error reading question from line: %s", err.Error())
			continue
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func questionFromString(line, separator string) (question.Question, error) {
	parts := strings.Split(line, separator)
	if len(parts) != 8 {
		return question.Question{}, errors.New("error in question line")
	}
	id, _ := strconv.Atoi(parts[0])
	category := parts[1]
	difficulty := parts[2]
	text := parts[3]
	answer := parts[4]
	return *question.New(id, category, difficulty, text, answer), nil
}
