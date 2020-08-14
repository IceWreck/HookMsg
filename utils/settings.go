package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	DeploymentName     string `json:"deployment_name"`
	DeploymentURL      string `json:"localhost:3333"`
	TelegramUserName   string `json:"telegram_user"`
	TelegramUserChatID int64  `json:"telegram_user_chat_id"`
	TelegramToken      string `json:"telegram_token"`
	SMTPEmail          string
	SMTPPort           int
}

var Config Settings = loadSettings()

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
