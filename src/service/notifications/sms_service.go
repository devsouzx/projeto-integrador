package notifications

import (
	"fmt"

	config "github.com/devsouzx/projeto-integrador/src/configuration"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSService struct {
	config *config.TwilioConfig
}

func NewSMSService() *SMSService {
	return &SMSService{
		config: config.NewTwilioConfig(),
	}
}

func (s *SMSService) SendSMS(to, message string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: s.config.AccountSID,
		Password: s.config.AuthToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(s.config.FromNumber)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("erro ao enviar SMS: %v", err)
	}

	 fmt.Printf("Resposta Twilio: %+v\n", *resp)
	return nil
}