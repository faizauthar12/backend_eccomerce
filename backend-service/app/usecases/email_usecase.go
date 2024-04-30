package usecases

// import (
// 	"context"
// 	"crypto/tls"
// 	"embed"
// 	"fmt"
// 	"net"
// 	"net/smtp"

// 	"github.com/faizauthar12/eccomerce/backend-service/app/constants"
// )

// var (
// 	//go:embed templates
// 	templates embed.FS
// )

// type IEmailUseCase interface {
// 	SendWelcomeEmail()
// }

// type EmailUseCase struct {
// 	ctx context.Context
// }

// func NewEmailUseCase(
// 	ctx context.Context,
// ) IEmailUseCase {
// 	return &EmailUseCase{
// 		ctx: ctx,
// 	}
// }

// func (u *EmailUseCase) sendEmail(
// 	emailSender string,
// 	emailPassword string,
// 	emailRecipient string,
// 	emailSubject string,
// 	emailBody string,
// ) error {
// 	fmt.Println("email-gomod: sendMail")

// 	// Connect to the server, authenticate, set the sender and recipient,
// 	// and send the email all in one step.
// 	msg := []byte("To: " + emailRecipient + "\r\n" +
// 		"Subject: " + emailSubject + "\r\n" +
// 		"Content-Type: text/html\r\n" +
// 		"\r\n" +
// 		emailBody + "\r\n",
// 	)

// 	// Set up authentication information.
// 	auth := smtp.PlainAuth("", emailSender, emailPassword, constants.CONFIG_SMTP_HOST)
// 	smtpAddr := fmt.Sprintf("%s:%d", constants.CONFIG_SMTP_HOST, constants.CONFIG_SMTP_PORT)

// 	// Connect to the server
// 	connection, errorConnection := net.Dial("tcp", smtpAddr)
// 	if errorConnection != nil {
// 		return errorConnection
// 	}

// 	// setup client
// 	client, errorClient := smtp.NewClient(connection, constants.CONFIG_SMTP_HOST)
// 	if errorClient != nil {
// 		fmt.Println("Error Client")
// 		return errorClient
// 	}

// 	// Start TLS
// 	tlsConfig := &tls.Config{
// 		InsecureSkipVerify: false,
// 		ServerName:         constants.CONFIG_SMTP_HOST,
// 	}

// 	if err := client.StartTLS(tlsConfig); err != nil {
// 		fmt.Println("Error Start TLS")
// 		return err
// 	}

// 	// authenticate
// 	errorAuthenticate := client.Auth(auth)
// 	if errorAuthenticate != nil {
// 		fmt.Println("Error Authenticate")
// 		return errorAuthenticate
// 	}

// 	// set the sender
// 	errorSetSender := client.Mail(constants.MAIL_SENDER)
// 	if errorSetSender != nil {
// 		fmt.Println("Error Set Sender")
// 		return errorSetSender
// 	}

// 	// set the recipient
// 	errorSetRecipient := client.Rcpt(emailRecipient)
// 	if errorSetRecipient != nil {
// 		fmt.Println("Error Set Recipient")
// 		return errorSetRecipient
// 	}

// 	// write the email body to the email
// 	wc, errorWC := client.Data()
// 	if errorWC != nil {
// 		return errorWC
// 	}

// 	defer wc.Close()

// 	_, errorWrite := wc.Write(msg)
// 	return errorWrite
// }
