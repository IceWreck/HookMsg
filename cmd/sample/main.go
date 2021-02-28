// Sample file for testing stupid shit

package main

import (
	"fmt"

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
		fmt.Println("err logging in", err)
	}
	client.SetCredentials(resp.UserID, resp.AccessToken)
	client.Store.SaveFilterID(config.Config.MatrixUserName, config.Config.MatrixTerminalFilter)
	syncer := client.Syncer.(*matrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", func(ev *matrix.Event) {
		body, _ := ev.Body()
		// if ok && ev.RoomID == "!E6WUDKPsMZL88Oth:chat.abifog.com" {
		//
		// }

		fmt.Println(body)
	})
	if err := client.Sync(); err != nil {
		fmt.Println("Sync() returned ", err)
	}
}
