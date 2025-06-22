package notifications

import "fmt"

type NotificationService struct {
	smsService *SMSService
}

func NewNotificationService(sms *SMSService) *NotificationService {
	return &NotificationService{smsService: sms}
}

func (n *NotificationService) SendExamNotification(phone, patientName, returnDate string) error {
	message := fmt.Sprintf(
		"Olá %s, seu exame foi realizado. Retorne em %s para resultados. Dúvidas: sobrevidas.ccu@gmail.com",
		patientName,
		returnDate,
	)

	phone = "+55" + phone

	return n.smsService.SendSMS(phone, message)
}