package actions

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"

	"github.com/IceWreck/HookMsg/config"
)

// Script -
type Script struct {
	Shell    string `json:"shell"`
	Location string `json:"location"`
	Endpoint string `json:"endpoint"`
	Secret   string `json:"secret"`
}

// RunScript executes script of type Script in defined shell
func RunScript(app *config.Application, endpoint string, secret string, webhookData map[string]interface{}) {

	// Read enabled_scripts.json
	var scripts []Script
	jsonFile, err := os.Open(app.Config.ScriptsConfig)
	if err != nil {
		app.Logger.Warn().Msg("Script config file not found")
		return
	} else {
		app.Logger.Info().Msg("Loaded script config file")
		defer jsonFile.Close()
	}

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &scripts)

	// iterate over scripts in json file and find one with matching credentials
	for _, s := range scripts {
		app.Logger.Debug().Str("current_script", s.Endpoint).Msg("Finding script....")
		if s.Endpoint == endpoint {
			app.Logger.Debug().Str("current_script", s.Endpoint).Msg("Found script")
			// now check if secret matches
			// you need not send a secret if secret is empty
			if s.Secret != secret {
				app.Logger.Warn().Str("current_script", s.Endpoint).Msg("Unauthorized")
				break
			}

			// Everything checks out. Execute script.
			webhookDataJSON, _ := json.Marshal(webhookData)

			//app.Logger.Debug().Interface("webhookData", string(webhookDataJSON)).Msg("")

			out, _ := exec.Command(s.Shell, s.Location, string(webhookDataJSON)).CombinedOutput()
			app.Logger.Debug().Str("current_script", s.Endpoint).Str("output", string(out)).Msg("")
			break
		}

	}

}
