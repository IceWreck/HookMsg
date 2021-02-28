package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config - settings which you load from your JSON
var Config Settings = loadSettings()

// Settings - struct to define settings
type Settings struct {
	DeploymentName       string                   `yaml:"deployment_name"`
	DeploymentPort       int                      `yaml:"deployment_port"`
	DeploymentURL        string                   `yaml:"deployment_url"`
	TelegramUserName     string                   `yaml:"tg_user"`
	TelegramUserChatID   int64                    `yaml:"tg_user_chat_id"`
	TelegramToken        string                   `yaml:"tg_token"`
	TelegramKey          []string                 `yaml:"tg_key"`
	MatrixUserName       string                   `yaml:"matrix_user"`
	MatrixPassword       string                   `yaml:"matrix_password"`
	MatrixHomeserver     string                   `yaml:"matrix_homeserver"`
	MatrixDeviceID       string                   `yaml:"matrix_deviceid"`
	MatrixChannels       map[string]MatrixChannel `yaml:"matrix_channels"`
	MatrixTerminal       string                   `yaml:"matrix_terminal"`
	MatrixTerminalUser   string                   `yaml:"matrix_terminal_user"`
	MatrixTerminalFilter string                   `yaml:"matrix_terminal_filter"`
	ScriptsConfig        string                   `yaml:"scripts_config"`
}

// MatrixChannel has matrix room id and hookmsg API key for that channel
// the string in the map in MatrixChannels is the channel short name for HookMsg
type MatrixChannel struct {
	ID  string `yaml:"id"`
	Key string `yaml:"key"`
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
	err = yaml.Unmarshal(byteValue, &settings)
	if err != nil {
		log.Fatal("Error parsing config.yaml")
	}
	return settings
}
