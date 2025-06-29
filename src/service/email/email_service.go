package email

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendRecoveryEmail(email, code string) error
	SendVerificationEmail(email, name, verificationLink, token string) error
	SendAgendamentoCreatedEmail(email string, agendamento *model.Agendamento) error
	SendAgendamentoConfirmedEmail(email string, agendamento *model.Agendamento) error
	SendAgendamentoCanceledEmail(email string, agendamento *model.Agendamento) error
	SendAgendamentoCompletedEmail(email string, agendamento *model.Agendamento) error
}

type EmailJob struct {
	Email            string
	Code             string
	Name             string
	VerificationLink string
	Token            string
	EmailType        string
	Agendamento      *model.Agendamento
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

		case "agendamento_created":
			m.SetHeader("Subject", "Recebemos seu agendamento!")
			loc, _ := time.LoadLocation("America/Sao_Paulo")
			dataHora := job.Agendamento.DataHora.In(loc).Add(24 * time.Hour).Format("02/01/2006 às 15:04")

			body := fmt.Sprintf(`
                <html>
                <body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
                    <div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px;">
                        <div style="text-align: center; margin-bottom: 20px;">
                            <h2 style="color: #9b59b6;">Agendamento Recebido</h2>
                            <div style="background-color: #f8f9fa; padding: 15px; border-radius: 6px; margin: 20px 0;">
                                <p style="margin: 0; font-size: 16px;">Aguarde a confirmação do profissional</p>
                            </div>
                        </div>
                        
                        <div style="background-color: #f8f9fa; padding: 15px; border-radius: 6px; margin-bottom: 20px;">
                            <h3 style="color: #9b59b6; margin-top: 0;">Detalhes do Agendamento</h3>
                            <table style="width: 100%%;">
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Data/Hora:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Tipo:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Status:</td>
                                    <td style="padding: 5px 0;"><span style="color: #ffc107; font-weight: bold;">%s</span></td>
                                </tr>
                            </table>
                        </div>
                        
                        <p style="text-align: center; color: #6c757d;">Você receberá uma nova notificação quando seu agendamento for confirmado.</p>
                        
                        <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #e0e0e0; text-align: center; color: #6c757d; font-size: 14px;">
                            <p>Em caso de dúvidas, entre em contato conosco.</p>
                        </div>
                    </div>
                </body>
                </html>
            `, dataHora, job.Agendamento.TipoConsulta, job.Agendamento.Status)
			m.SetBody("text/html", body)

		case "agendamento_confirmed":
			m.SetHeader("Subject", "Seu agendamento foi confirmado!")
			loc, _ := time.LoadLocation("America/Sao_Paulo")
			dataHora := job.Agendamento.DataHora.In(loc).Add(24 * time.Hour).Format("02/01/2006 às 15:04")

			body := fmt.Sprintf(`
                <html>
                <body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
                    <div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px;">
                        <div style="text-align: center; margin-bottom: 20px;">
                            <h2 style="color: #28a745;">Agendamento Confirmado!</h2>
                            <div style="background-color: #e6ffed; padding: 15px; border-radius: 6px; margin: 20px 0;">
                                <p style="margin: 0; font-size: 16px;">Seu agendamento foi confirmado pelo profissional</p>
                            </div>
                        </div>
                        
                        <div style="background-color: #f8f9fa; padding: 15px; border-radius: 6px; margin-bottom: 20px;">
                            <h3 style="color: #28a745; margin-top: 0;">Detalhes do Agendamento</h3>
                            <table style="width: 100%%;">
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Data/Hora:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Tipo:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Status:</td>
                                    <td style="padding: 5px 0;"><span style="color: #28a745; font-weight: bold;">%s</span></td>
                                </tr>
                            </table>
                        </div>
                        
                        <div style="background-color: #fff3cd; padding: 15px; border-radius: 6px; margin-bottom: 20px;">
                            <h4 style="margin-top: 0; color: #856404;">Lembrete</h4>
                            <p>Por favor, chegue com 15 minutos de antecedência e traga seus documentos.</p>
                        </div>
                        
                        <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #e0e0e0; text-align: center; color: #6c757d; font-size: 14px;">
                            <p>Em caso de dúvidas ou necessidade de reagendamento, entre em contato conosco.</p>
                        </div>
                    </div>
                </body>
                </html>
            `, dataHora, job.Agendamento.TipoConsulta, job.Agendamento.Status)
			m.SetBody("text/html", body)

		case "agendamento_canceled":
			m.SetHeader("Subject", "Seu agendamento foi cancelado")
			loc, _ := time.LoadLocation("America/Sao_Paulo")
			dataHora := job.Agendamento.DataHora.In(loc).Add(24 * time.Hour).Format("02/01/2006 às 15:04")

			body := fmt.Sprintf(`
                <html>
                <body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
                    <div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px;">
                        <div style="text-align: center; margin-bottom: 20px;">
                            <h2 style="color: #dc3545;">Agendamento Cancelado</h2>
                            <div style="background-color: #f8d7da; padding: 15px; border-radius: 6px; margin: 20px 0;">
                                <p style="margin: 0; font-size: 16px;">Seu agendamento foi cancelado</p>
                            </div>
                        </div>
                        
                        <div style="background-color: #f8f9fa; padding: 15px; border-radius: 6px; margin-bottom: 20px;">
                            <h3 style="color: #dc3545; margin-top: 0;">Detalhes do Agendamento</h3>
                            <table style="width: 100%%;">
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Data/Hora:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Tipo:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Status:</td>
                                    <td style="padding: 5px 0;"><span style="color: #dc3545; font-weight: bold;">%s</span></td>
                                </tr>
                            </table>
                        </div>
                        
                        <div style="text-align: center; margin: 25px 0;">
                            <a href="[LINK_PARA_NOVO_AGENDAMENTO]" style="background: #4a6baf; color: white; padding: 12px 24px; 
                               text-decoration: none; border-radius: 4px; font-weight: bold;">
                                Agendar Novamente
                            </a>
                        </div>
                        
                        <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #e0e0e0; text-align: center; color: #6c757d; font-size: 14px;">
                            <p>Sentimos muito pelo inconveniente. Para reagendar ou esclarecer dúvidas, entre em contato conosco.</p>
                        </div>
                    </div>
                </body>
                </html>
            `, dataHora, job.Agendamento.TipoConsulta, job.Agendamento.Status)
			m.SetBody("text/html", body)

		case "agendamento_completed":
			m.SetHeader("Subject", "Consulta concluída - Avalie nosso serviço")
			loc, _ := time.LoadLocation("America/Sao_Paulo")
			dataHora := job.Agendamento.DataHora.In(loc).Add(24 * time.Hour).Format("02/01/2006 às 15:04")

			body := fmt.Sprintf(`
                <html>
                <body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
                    <div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px;">
                        <div style="text-align: center; margin-bottom: 20px;">
                            <h2 style="color: #6f42c1;">Consulta Concluída</h2>
                            <div style="background-color: #e9ecef; padding: 15px; border-radius: 6px; margin: 20px 0;">
                                <p style="margin: 0; font-size: 16px;">Obrigado por utilizar nossos serviços!</p>
                            </div>
                        </div>
                        
                        <div style="background-color: #f8f9fa; padding: 15px; border-radius: 6px; margin-bottom: 20px;">
                            <h3 style="color: #6f42c1; margin-top: 0;">Detalhes da Consulta</h3>
                            <table style="width: 100%%;">
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Data/Hora:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Tipo:</td>
                                    <td style="padding: 5px 0;">%s</td>
                                </tr>
                                <tr>
                                    <td style="width: 120px; padding: 5px 0; font-weight: bold;">Status:</td>
                                    <td style="padding: 5px 0;"><span style="color: #6f42c1; font-weight: bold;">%s</span></td>
                                </tr>
                            </table>
                        </div>
                        
                        <div style="text-align: center; margin: 25px 0;">
                            <p style="font-size: 16px;">Avalie sua experiência conosco:</p>
                            <div>
                                <a href="[LINK_AVALIACAO_5]" style="display: inline-block; margin: 0 5px; color: #ffc107; font-size: 24px; text-decoration: none;">★</a>
                                <a href="[LINK_AVALIACAO_4]" style="display: inline-block; margin: 0 5px; color: #ffc107; font-size: 24px; text-decoration: none;">★</a>
                                <a href="[LINK_AVALIACAO_3]" style="display: inline-block; margin: 0 5px; color: #ffc107; font-size: 24px; text-decoration: none;">★</a>
                                <a href="[LINK_AVALIACAO_2]" style="display: inline-block; margin: 0 5px; color: #ffc107; font-size: 24px; text-decoration: none;">★</a>
                                <a href="[LINK_AVALIACAO_1]" style="display: inline-block; margin: 0 5px; color: #ffc107; font-size: 24px; text-decoration: none;">★</a>
                            </div>
                        </div>
                        
                        <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #e0e0e0; text-align: center; color: #6c757d; font-size: 14px;">
                            <p>Seu feedback é muito importante para melhorarmos nossos serviços.</p>
                        </div>
                    </div>
                </body>
                </html>
            `, dataHora, job.Agendamento.TipoConsulta, job.Agendamento.Status)
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
		Email:            email,
		Name:             name,
		VerificationLink: verificationLink,
		Token:            token,
		EmailType:        "verification",
	}
	return nil
}

func (es *emailService) SendAgendamentoCreatedEmail(email string, agendamento *model.Agendamento) error {
	es.jobQueue <- EmailJob{
		Email:       email,
		Agendamento: agendamento,
		EmailType:   "agendamento_created",
	}
	return nil
}

func (es *emailService) SendAgendamentoConfirmedEmail(email string, agendamento *model.Agendamento) error {
	es.jobQueue <- EmailJob{
		Email:       email,
		Agendamento: agendamento,
		EmailType:   "agendamento_confirmed",
	}
	return nil
}

func (es *emailService) SendAgendamentoCanceledEmail(email string, agendamento *model.Agendamento) error {
	es.jobQueue <- EmailJob{
		Email:       email,
		Agendamento: agendamento,
		EmailType:   "agendamento_canceled",
	}
	return nil
}

func (es *emailService) SendAgendamentoCompletedEmail(email string, agendamento *model.Agendamento) error {
	es.jobQueue <- EmailJob{
		Email:       email,
		Agendamento: agendamento,
		EmailType:   "agendamento_completed",
	}
	return nil
}
