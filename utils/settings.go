package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config - settings which you load from your JSON
var Config Settings = loadSettings()

// Settings - struct to define settings
type Settings struct {
	DeploymentName      string   `yaml:"deployment_name"`
	DeploymentPort      int      `yaml:"deployment_port"`
	DeploymentURL       string   `yaml:"deployment_url"`
	TelegramUserName    string   `yaml:"tg_user"`
	TelegramUserChatID  int64    `yaml:"tg_user_chat_id"`
	TelegramToken       string   `yaml:"tg_token"`
	TelegramWebhookAuth []string `yaml:"tg_webhook_auth"`
	ScriptsConfig       string   `yaml:"scripts_config"`
}

func loadSettings() Settings {
	var settings Settings
	yamlFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("Config file not found.")
	}
	defer yamlFile.Close()
	log.Println("Loaded config.yaml")
	byteValue, _ := ioutil.ReadAll(yamlFile)
	yaml.Unmarshal(byteValue, &settings)
	fmt.Println(settings)
	return settings
}
