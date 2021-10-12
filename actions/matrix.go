package actions

import (
	"bytes"
	"time"

	"github.com/IceWreck/HookMsg/config"
	"github.com/matrix-org/gomatrix"
	"github.com/yuin/goldmark"
)

func MatrixClientInit(app *config.Application) *gomatrix.Client {
	// login initially
	app.Logger.Info().Msg("Logging into Matrix")
	c, _ := gomatrix.NewClient(app.Config.MatrixHomeserver, "", "")
	clientLogin(app, c)

	// start ticker to re-login every week
	ticker := time.NewTicker(7 * 24 * time.Hour)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				app.Logger.Info().Msg("Attempting scheduled matrix relogin")
				clientLogin(app, c)
			}
		}
	}()

	return c
}

func clientLogin(app *config.Application, c *gomatrix.Client) {
	// TODO: while probably not required but put this in a mutex just in case
	resp, err := c.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     app.Config.MatrixUserName,
		Password: app.Config.MatrixPassword,
		DeviceID: app.Config.MatrixDeviceID,
	})
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error logging in to matrix")
	} else {
		app.Logger.Info().Msg("Logged into matrix")
	}
	c.SetCredentials(resp.UserID, resp.AccessToken)
}

// SendMatrixText - send text message on given matrix channel
func SendMatrixText(app *config.Application, id string, body string) {
	// user will send markdown
	// body will remail markdown
	// formattedBody should be converted to html
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(body), &buf); err != nil {
		app.Logger.Error().Err(err).Msg("Error converting markdown to html")
		return
	}
	_, err := app.MatrixClient.SendFormattedText(id, body, buf.String())
	if err != nil {
		app.Logger.Error().Err(err).Msg("")
		// retry logging in
		clientLogin(app, app.MatrixClient)
		// retry sending
		app.MatrixClient.SendFormattedText(id, body, body)
	}
}
