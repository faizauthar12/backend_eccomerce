package models

type WelcomeEmail struct {
	BaseURL        string
	EmailSubject   string
	EmailBody      string
	RecipientEmail string
}
