package service

import (
	"data-pipeline/config"
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func SendMail(from string, to string, subject string, content string) error {
	client := mailjet.NewMailjetClient(config.MailJetKey, config.MailJetSecret)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: from,
				Name:  from,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
				},
			},
			Subject:  subject,
			HTMLPart: fmt.Sprintf("<h3>%s</h3>", content),
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := client.SendMailV31(&messages)
	if err != nil {
		return err
	}

	return nil
}
