package main

import "fmt"

type Notification struct {
	Recepient string
	Subject   string
	Body      string
	Channel   string
	IsUrgent  bool
}

func main() {
	unurgent := Notification{
		Recepient: "mbrandosha@gmail.com",
		Subject:   "Application for status confirmation.",
		Body:      "Xin Chao. Your request was confirmed!",
		Channel:   "email",
		IsUrgent:  false,
	}

	adminAlert := Notification{
		Recepient: "@meteor_jam",
		Subject:   "Deploy Status Information",
		Body:      "Everything is good. Deployment completed successfully!",
		Channel:   "telegram",
		IsUrgent:  true,
	}

	noRecepient := Notification{
		Recepient: "",
		Subject:   "Subject.",
		Body:      "Body text.",
		Channel:   "telegram",
		IsUrgent:  true,
	}

	runCase(unurgent)
	runCase(adminAlert)
	runCase(noRecepient)
}

func (n Notification) Send() (string, string) {
	message := n.Format()
	status := "queued"
	return message, status
}

func (n Notification) Format() string {
	return fmt.Sprintf(
		"Subject: %s Body: %s",
		n.Subject,
		n.Body,
	)
}

func (n Notification) ValidateRecepint() error {
	if n.Recepient == "" {
		return fmt.Errorf("Error sending notification.")
	}

	return nil
}

func runCase(n Notification) {
	if err := n.ValidateRecepint(); err != nil {
		fmt.Println("Oops!", err.Error())
		return
	}
	message, status := n.Send()

	fmt.Println("Notification:")
	fmt.Println(message)
	fmt.Println("Status:", status)
	fmt.Println()
}
