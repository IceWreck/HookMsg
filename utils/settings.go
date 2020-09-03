package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config - settings which you load from your JSON
var Config Settings = loadSettings()

// Settings - struct to define settings
type Settings struct {
	DeploymentName      string   `json:"deployment_name"`
	DeploymentPort      int      `json:"deployment_port"`
	DeploymentURL       string   `json:"deployment_url"`
	TelegramUserName    string   `json:"tg_user"`
	TelegramUserChatID  int64    `json:"tg_user_chat_id"`
	TelegramToken       string   `json:"tg_token"`
	TelegramWebhookAuth []string `json:"tg_webhook_auth"`
	ScriptsConfig       string   `json:"scripts_config"`
	SMTPEmail           string
	SMTPPort            int
}

func loadSettings() Settings {
	var config Settings
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Config file not found.")
	}
	defer jsonFile.Close()
	log.Println("Loaded config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)
	return config
}
