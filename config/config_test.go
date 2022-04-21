package config

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	config, err := FromEnvFile("../local.env")
	if err != nil {
		t.Error(err)
	}
	if config.FirebaseConfig == "" {
		t.Error("error loading firebase config")
	}
}
