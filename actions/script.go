package actions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Script -
type Script struct {
	Shell    string `json:"shell"`
	Location string `json:"location"`
	Endpoint string `json:"endpoint"`
	Secret   string `json:"secret"`
}

// RunScript executes script of type Script in defined shell
func RunScript(formResults map[string]string) {

	// Read enabled_scripts.json
	var scripts []Script
	jsonFile, err := os.Open("enabled_scripts.json")
	if err != nil {
		log.Println("enabled_scripts.json not found.")
	} else {
		log.Println("Loaded enabled_scripts.json")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &scripts)

	// iterate over scripts in json file and find one with matching credentials
	for _, s := range scripts {
		log.Println(s)
		if s.Endpoint == formResults["endpoint"] {
			// now check if secret matches
			// you need not send a secret if secret is empty
			if s.Secret != formResults["secret"] {
				log.Println("Secrets dont match.")
				break
			}
			out, _ := exec.Command(s.Shell, s.Location).CombinedOutput()
			log.Print(string(out))
			break
		}

	}

}
