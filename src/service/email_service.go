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

type EmailJob struct {
    Email string
    Code  string
}

type emailService struct {
    dialer   *gomail.Dialer
    jobQueue chan EmailJob
}

func NewEmailService() EmailService {
    smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
    
    es := &emailService{
        dialer: gomail.NewDialer(
            os.Getenv("SMTP_HOST"),
            smtpPort,
            os.Getenv("SMTP_USER"),
            os.Getenv("SMTP_PASS"),
        ),
        jobQueue: make(chan EmailJob, 100),
    }

    for i := 0; i < 5; i++ {
        go es.emailWorker()
    }
    
    return es
}

func (es *emailService) emailWorker() {
    for job := range es.jobQueue {
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
        `, job.Code)
        
        m.SetBody("text/html", body)
        
        if err := es.dialer.DialAndSend(m); err != nil {
            fmt.Printf("Erro ao enviar email: %v\n", err)
        }
    }
}

func (es *emailService) SendRecoveryEmail(email, code string) error {
    es.jobQueue <- EmailJob{Email: email, Code: code}
    return nil
}