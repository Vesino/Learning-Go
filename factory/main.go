package main

import "fmt"

// SMS, EMAIL only

type INotificationFactory interface {
	// As it is an interface we could apply a kind of poliformis
	// The following methods are REQUIRED for every instance of this interface
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

type SMSNotificationSender struct {
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct {
}

type EmailNotificationSender struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "EMAIL" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type allowed")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)
}
