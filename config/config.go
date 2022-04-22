package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

const (
	fbTypeKey        string = "FB_TYPE"
	fbProjectIDKey   string = "FB_PROJECT_ID"
	fbPrivateKeyID   string = "FB_PRIVATE_KEY_ID"
	fbPrivateKeyKey  string = "FB_PRIVATE_KEY"
	fbClientIDKey    string = "FB_CLIENT_ID"
	fbClientEmailKey string = "FB_CLIENT_EMAIL"
	fbTokenURIKey    string = "FB_TOKEN_URI"
	fbAuthURIKey     string = "FB_AUTH_URI"
	fbAuthCertURL    string = "FB_AUTH_PROVIDER_CERT_URL"
	fbClientCertURL  string = "FB_CLIENT_CERT_URL"

	firebaseConfigKey    string = "FIREBASE_CONFIG"
	firebaseConfigB64Key string = "FIREBASE_CONFIG_B64"
)

type (
	Config struct {
		fbConfig firebaseConfig
	}
	firebaseConfig struct {
		AccountType   string `json:"type"`
		ProjectID     string `json:"project_id"`
		PrivateKeyID  string `json:"private_key_id"`
		PrivateKey    string `json:"private_key"`
		ClientEmail   string `json:"client_email"`
		ClientID      string `json:"client_id"`
		AuthURI       string `json:"auth_uri"`
		TokenURI      string `json:"token_uri"`
		AuthCertURL   string `json:"auth_provider_x509_cert_url"`
		ClientCertURL string `json:"client_x509_cert_url"`
	}
)

// New creates and initializes config from the current environment
func New() (Config, error) {
	// load all environment variables from current ENV
	fbConfig := firebaseConfig{
		AccountType:  os.Getenv(fbTypeKey),
		ProjectID:    os.Getenv(fbProjectIDKey),
		PrivateKeyID: os.Getenv(fbPrivateKeyID),
		PrivateKey:   os.Getenv(fbPrivateKeyKey),
		ClientID:     os.Getenv(fbClientIDKey),
		ClientEmail:  os.Getenv(fbClientEmailKey),
		//AuthURI:       os.Getenv(fbAuthURIKey),
		//TokenURI:      os.Getenv(fbTokenURIKey),
		//AuthCertURL:   os.Getenv(fbAuthCertURL),
		//ClientCertURL: os.Getenv(fbClientCertURL),
	}
	return Config{fbConfig: fbConfig}, nil
}

// FromEnvFile loads config with environment variables from a specific ENV file
func FromEnvFile(filePath string) (Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return Config{}, err
	}
	return New()
}

func (config Config) Stringify() string {
	bytes, _ := json.Marshal(config.fbConfig)
	return string(bytes)
}

func (config Config) Bytes() []byte {
	bytes, _ := json.Marshal(config.fbConfig)
	return bytes
}
