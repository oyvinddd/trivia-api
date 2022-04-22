package importer

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/oyvinddd/trivia-api/config"
	"os"
	"testing"
)

func TestImportQuestions(t *testing.T) {
	godotenv.Load("../local.env")
	cfg := config.Firebase(
		os.Getenv("FB_TYPE"),
		os.Getenv("FB_PROJECT_ID"),
		os.Getenv("FB_PRIVATE_KEY_ID"),
		os.Getenv("FB_PRIVATE_KEY"),
		os.Getenv("FB_CLIENT_EMAIL"),
		os.Getenv("FB_CLIENT_ID"))
	importer := NewFirebaseImporter(context.TODO(), cfg)
	err := importer.ImportAvailableQuestions()
	if err != nil {
		t.Error(err)
	}
}
