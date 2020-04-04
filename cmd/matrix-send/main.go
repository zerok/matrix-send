package main

import (
	"io/ioutil"
	"os"

	"github.com/matrix-org/gomatrix"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	var roomID string
	pflag.StringVar(&roomID, "room", "", "Room ID")
	pflag.Parse()

	homeserver := os.Getenv("MATRIX_HOMESERVER_URL")
	username := os.Getenv("MATRIX_USER")
	password := os.Getenv("MATRIX_PASSWORD")
	if roomID == "" {
		logger.Fatal().Msg("Please set --room.")
	}
	if username == "" {
		logger.Fatal().Msg("Please set $MATRIX_USER.")
	}
	if password == "" {
		logger.Fatal().Msg("Please set $MATRIX_PASSWORD.")
	}
	if homeserver == "" {
		logger.Fatal().Msg("Please set $MATRIX_HOMESERVER_URL.")
	}
	client, err := gomatrix.NewClient(homeserver, "", "")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create new client.")
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to read from stdin.")
	}

	resp, err := client.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     username,
		Password: password,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to log in.")
	}
	client.SetCredentials(resp.UserID, resp.AccessToken)
	if _, err := client.SendText(roomID, string(data)); err != nil {
		logger.Fatal().Err(err).Msg("Failed to send text.")
	}
}
