package config

import "os"

type TwilioConfig struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

func NewTwilioConfig() *TwilioConfig {
	return &TwilioConfig{
		AccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		FromNumber: os.Getenv("TWILIO_PHONE"),
	}
}