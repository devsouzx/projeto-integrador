package email

import (
	"fmt"
	"os"
	"strconv"
    
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendRecoveryEmail(email, code string) error
	SendVerificationEmail(email, name, verificationLink, token string) error
}

type EmailJob struct {
    Email           string
    Code            string
    Name            string
    VerificationLink string
    Token           string
    EmailType       string
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
        m.SetHeader("To", job.Email)
        
        switch job.EmailType {
        case "recovery":
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
            
        case "verification":
            m.SetHeader("Subject", "Ativação de Conta - Sistema Médico")
            body := fmt.Sprintf(`
                <html>
                <body style="font-family: Arial, sans-serif; line-height: 1.6;">
                    <div style="max-width: 600px; margin: 0 auto;">
                        <h2 style="color: #9b59b6;">Ativação de Conta</h2>
                        <p>Olá <strong>%s</strong>,</p>
                        <p>Obrigado por se cadastrar em nosso sistema médico! Para ativar sua conta, clique no link abaixo:</p>
                        
                        <div style="margin: 25px 0;">
                            <a href="%s" style="background: #9b59b6; color: white; padding: 12px 24px; 
                               text-decoration: none; border-radius: 4px; font-weight: bold;">
                                Ativar Minha Conta
                            </a>
                        </div>
                        
                        <p>Ou copie este link em seu navegador:<br>
                        <code style="word-break: break-all;">%s</code></p>
                        
                        <p style="margin-top: 20px; font-size: 0.9em; color: #9b59b6;">
                            <strong>Token de verificação:</strong> %s<br>
                            Este link expirará em 1 hora.
                        </p>
                        
                        <p style="font-size: 0.9em; color: #9b59b6;">
                            Caso não tenha criado uma conta conosco, ignore este e-mail.
                        </p>
                    </div>
                </body>
                </html>
            `, job.Name, job.VerificationLink, job.VerificationLink, job.Token)
            m.SetBody("text/html", body)
        }

        if err := es.dialer.DialAndSend(m); err != nil {
            fmt.Printf("Erro ao enviar email (%s): %v\n", job.EmailType, err)
        }
    }
}

func (es *emailService) SendRecoveryEmail(email, code string) error {
    es.jobQueue <- EmailJob{
        Email:     email,
        Code:      code,
        EmailType: "recovery",
    }
    return nil
}

func (es *emailService) SendVerificationEmail(email, name, verificationLink, token string) error {
    es.jobQueue <- EmailJob{
        Email:           email,
        Name:           name,
        VerificationLink: verificationLink,
        Token:          token,
        EmailType:      "verification",
    }
    return nil
}