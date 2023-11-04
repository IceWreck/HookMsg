package actions

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

// Script -
type Script struct {
	Shell    string `json:"shell"`
	Location string `json:"location"`
	Endpoint string `json:"endpoint"`
	Secret   string `json:"secret"`
}

// RunScript executes script of type Script in defined shell
func (svc *Service) RunScript(endpoint string, secret string, webhookData map[string]interface{}) {

	// Read enabled_scripts.json
	var scripts []Script
	jsonFile, err := os.Open(svc.config.ScriptsConfig)
	if err != nil {
		log.Warn().Msg("Script config file not found")
		return
	} else {
		log.Info().Msg("Loaded script config file")
		defer jsonFile.Close()
	}

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &scripts)

	// iterate over scripts in json file and find one with matching credentials
	for _, s := range scripts {
		log.Debug().Str("current_script", s.Endpoint).Msg("Finding script....")
		if s.Endpoint == endpoint {
			log.Debug().Str("current_script", s.Endpoint).Msg("Found script")
			// now check if secret matches
			// you need not send a secret if secret is empty
			if s.Secret != secret {
				log.Warn().Str("current_script", s.Endpoint).Msg("Unauthorized")
				break
			}

			// Everything checks out. Execute script.
			webhookDataJSON, _ := json.Marshal(webhookData)

			//log.Debug().Interface("webhookData", string(webhookDataJSON)).Msg("")

			out, _ := exec.Command(s.Shell, s.Location, string(webhookDataJSON)).CombinedOutput()
			log.Debug().Str("current_script", s.Endpoint).Str("output", string(out)).Msg("")
			break
		}

	}
}
