package config

import (
	"encoding/json"
	"fmt"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type From struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type emailAddress struct {
	Email string `json:"email"`
}

type mailContent struct {
	From     From
	To       []emailAddress
	Subject  string
	Html     string
	Category *string
	Bcc      []emailAddress
	CC       []emailAddress
}

func setFrom(email string, name string) From {
	return From{
		Email: email,
		Name:  name,
	}
}

func SetEmailAddress(emails []string) []emailAddress {
	emailList := []emailAddress{}

	for _, email := range emails {
		emailList = append(emailList, emailAddress{
			Email: email,
		})
	}

	return emailList
}

func SetMailContent(
	to []emailAddress,
	subject string,
	html string,
	category *string,
	bcc []emailAddress,
	cc []emailAddress,
) *mailContent {
	return &mailContent{
		From:     setFrom(os.Getenv("MAILTRAP_EMAIL"), os.Getenv("MAILTRAP_NAME")),
		To:       to,
		Subject:  subject,
		Html:     html,
		Category: category,
		Bcc:      bcc,
		CC:       cc,
	}
}

func SendMailByMailjet() {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messageInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "me@gmail.com",
				Name:  "Me",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{

					Email: "you@concrete-corp.com",
					Name:  "You",
				},
			},
			Subject:  "My first Mailjet Email!",
			TextPart: "Greetings from Mailjet!",
		},
	}
	messages := &mailjet.MessagesV31{Info: messageInfo}
	res, err := mailjetClient.SendMailV31(messages)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}

func (mailContent *mailContent) SendMailWithMailtrap() bool {
	url := os.Getenv("MAILTRAP_HOST")
	method := "POST"

	from, _ := json.Marshal(mailContent.From)
	to, _ := json.Marshal(mailContent.To)
	subject, _ := json.Marshal(mailContent.Subject)
	html, _ := json.Marshal(mailContent.Html)
	category, _ := json.Marshal(mailContent.Category)

	payload := strings.NewReader(fmt.Sprintf(`{
		"from":%s,
		"to":%s,
		"subject":%s,
		"html":%s,
		"category":%s
		}`,
		from,
		to,
		subject,
		html,
		category))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Authorization", "Bearer "+os.Getenv("MAILTRAP_API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(body))

	return true
}
