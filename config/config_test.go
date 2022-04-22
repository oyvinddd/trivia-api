package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	_, err := FromEnvFile("../local.env")
	if err != nil {
		t.Error(err)
	}
}

func TestStringifyConfig(t *testing.T) {
	config, _ := FromEnvFile("../local.env")
	fmt.Println(config.Stringify())
}
