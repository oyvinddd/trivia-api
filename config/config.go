package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	firebaseConfigKey string = "FIREBASE_CONFIG"
)

type Config struct {
	FirebaseConfig string
}

// New creates and initializes config from the current environment
func New() Config {
	// load all environment variables from current ENV
	firebaseConfig := os.Getenv(firebaseConfigKey)

	config := Config{FirebaseConfig: firebaseConfig}
	return config
}

// FromEnvFile loads config with environment variables from a specific ENV file
func FromEnvFile(filePath string) (Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return Config{}, err
	}
	return New(), nil
}
