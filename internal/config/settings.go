package config

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// Settings - struct to define settings
type Config struct {
	DeploymentName       string                   `yaml:"deployment_name"`
	DeploymentPort       int                      `yaml:"deployment_port"`
	DeploymentURL        string                   `yaml:"deployment_url"`
	TelegramEnabled      bool                     `yaml:"tg_enabled"`
	TelegramUserName     string                   `yaml:"tg_user"`
	TelegramUserChatID   int64                    `yaml:"tg_user_chat_id"`
	TelegramToken        string                   `yaml:"tg_token"`
	TelegramKey          []string                 `yaml:"tg_key"`
	MatrixEnabled        bool                     `yaml:"matrix_enabled"`
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

func LoadConfig() Config {
	var settings Config
	yamlFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal().Msg("Config file not found")
	}
	defer yamlFile.Close()
	log.Info().Msg("Loaded config.yaml")
	byteValue, _ := io.ReadAll(yamlFile)
	err = yaml.Unmarshal(byteValue, &settings)
	if err != nil {
		log.Fatal().Msg("Error parsing config.yaml")
	}
	return settings
}
