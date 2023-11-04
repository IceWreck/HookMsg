package actions

import (
	"bytes"
	"time"

	"github.com/matrix-org/gomatrix"
	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
)

func (svc *Service) initMatrixClient() {
	// create client
	c, err := gomatrix.NewClient(svc.config.MatrixHomeserver, "", "")
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating matrix client")
	}
	svc.matrixClient = c

	// login initially
	svc.matrixClientLogin()

	// start ticker to re-login every week
	ticker := time.NewTicker(7 * 24 * time.Hour)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				log.Info().Msg("Attempting scheduled matrix relogin")
				svc.matrixClientLogin()
			}
		}
	}()
}

func (svc *Service) matrixClientLogin() {
	// TODO: while probably not required but put this in a mutex just in case
	log.Info().Msg("Logging into Matrix")
	resp, err := svc.matrixClient.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     svc.config.MatrixUserName,
		Password: svc.config.MatrixPassword,
		DeviceID: svc.config.MatrixDeviceID,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error logging in to matrix")
	}
	svc.matrixClient.SetCredentials(resp.UserID, resp.AccessToken)
	log.Info().Msg("Logged into matrix")
}

// SendMatrixText sends a text message on given matrix channel.
func (svc *Service) SendMatrixText(id string, body string) {
	if svc.matrixClient == nil {
		log.Error().Msg("Cannot send matrix text, client has not been initialized")
		return
	}

	// user will send markdown
	// body will remain markdown
	// formattedBody should be converted to html
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(body), &buf); err != nil {
		log.Error().Err(err).Msg("Error converting markdown to html")
		return
	}
	_, err := svc.matrixClient.SendFormattedText(id, body, buf.String())
	if err != nil {
		log.Error().Err(err).Msg("Error sending matrix text")
	}
}
