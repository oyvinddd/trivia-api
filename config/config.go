package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
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

// Firebase creates and initializes a Firebase based config from the current environment
func Firebase(accountType, projectID, privateKeyID, privateKey, clientEmail, clientID string) Config {
	config := Config{
		fbConfig: firebaseConfig{
			AccountType:  accountType,
			ProjectID:    projectID,
			PrivateKeyID: privateKeyID,
			PrivateKey:   privateKey,
			ClientEmail:  clientEmail,
			ClientID:     clientID,
		},
	}
	return config
}

// FromEnvFile loads config with environment variables from a specific ENV file
func FromEnvFile(filePath string) (Config, error) {
	if err := godotenv.Load(filePath); err != nil {
		return Config{}, err
	}
	return Config{}, nil
}

func (config Config) Stringify() string {
	bytes, _ := json.Marshal(config.fbConfig)
	return string(bytes)
}

func (config Config) Bytes() []byte {
	bytes, _ := json.Marshal(config.fbConfig)
	return bytes
}

func (config Config) PrintFirebaseConfig() {
	bytes, _ := json.Marshal(config.fbConfig)
	fmt.Println(string(bytes))
}
