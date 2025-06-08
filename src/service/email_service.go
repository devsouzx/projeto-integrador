package service

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService interface {
    SendRecoveryEmail(email, code string) error
}

type emailService struct {
    dialer *gomail.Dialer
}

func NewEmailService() EmailService {
    smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
    
    return &emailService{
        dialer: gomail.NewDialer(
            os.Getenv("SMTP_HOST"),
            smtpPort,
            os.Getenv("SMTP_USER"),
            os.Getenv("SMTP_PASS"),
        ),
    }
}

func (es *emailService) SendRecoveryEmail(email, code string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", os.Getenv("SMTP_FROM"))
    m.SetHeader("To", "joaoemanuel2215@gmail.com")
    m.SetHeader("Subject", "Código de Recuperação de Senha")
    
    body := fmt.Sprintf(`
        <html>
        <body>
            <h2>Recuperação de Senha</h2>
            <p>Você solicitou a recuperação de senha. Utilize o código abaixo para continuar:</p>
            <h3>%s</h3>
            <p>Este código é válido por 10 minutos.</p>
            <p>Caso não tenha solicitado esta alteração, ignore este e-mail.</p>
        </body>
        </html>
    `, code)
    
    m.SetBody("text/html", body)
    
    return es.dialer.DialAndSend(m)
}
