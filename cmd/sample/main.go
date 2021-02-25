// Sample file for testing stupid shit

package main

import (
	"log"

	"github.com/IceWreck/HookMsg/config"
	matrix "github.com/matrix-org/gomatrix"
)

func main() {
	client, _ := matrix.NewClient(config.Config.MatrixHomeserver, "", "")
	resp, err := client.Login(&matrix.ReqLogin{
		Type:     "m.login.password",
		User:     config.Config.MatrixUserName,
		Password: config.Config.MatrixPassword,
		DeviceID: config.Config.MatrixDeviceID,
	})
	if err != nil {
		log.Fatal("err logging in", err)
	}
	client.SetCredentials(resp.UserID, resp.AccessToken)
	client.SendText(config.Config.MatrixChannels["security-info"].ID, "lol u")

}
