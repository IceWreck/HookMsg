package actions

import (
	"log"
	"time"

	"github.com/IceWreck/HookMsg/config"
	matrix "github.com/matrix-org/gomatrix"
)

var client = clientInit()

func clientInit() *matrix.Client {
	// login initially
	c, _ := matrix.NewClient(config.Config.MatrixHomeserver, "", "")
	clientLogin(c)

	// start ticker to re-login every week
	ticker := time.NewTicker(7 * 24 * time.Hour)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				clientLogin(c)
			}
		}
	}()

	return c
}

func clientLogin(c *matrix.Client) {
	// TODO: while probably not required but put this in a mutex just in case
	resp, err := c.Login(&matrix.ReqLogin{
		Type:     "m.login.password",
		User:     config.Config.MatrixUserName,
		Password: config.Config.MatrixPassword,
		DeviceID: config.Config.MatrixDeviceID,
	})
	if err != nil {
		log.Println("Error logging in to matrix", err)
	} else {
		log.Println("Logged into matrix")
	}
	c.SetCredentials(resp.UserID, resp.AccessToken)
}

// SendMatrixText - send text message on given matrix channel
func SendMatrixText(id string, body string) {
	_, err := client.SendText(id, body)
	if err != nil {
		log.Println(err)
		// retry logging in
		clientLogin(client)
		// retry sending
		client.SendText(id, body)
	}
}
