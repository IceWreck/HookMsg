// This is required because default sync sends all events
// But we want events from a single room only (the server terminal one)
// So we create a filter and upload it to the server to filter ID once
// We wont ever create it again, just use the existing one

// Sample file for testing stupid shit

package main

import (
	"encoding/json"
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

	newFilter := matrix.DefaultFilter()
	newFilter.Room.Rooms = append(newFilter.Room.Rooms, config.Config.MatrixTerminal)
	newFilter.Room.Timeline.Senders = append(newFilter.Room.Timeline.Senders, config.Config.MatrixTerminalUser)
	filterJSON, _ := json.Marshal(newFilter)
	fmt.Println(string(filterJSON))
	filterResp, err := client.CreateFilter(filterJSON)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your FilterID is ", filterResp.FilterID)
	}
}
