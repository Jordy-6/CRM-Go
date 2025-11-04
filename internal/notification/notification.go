package notification

import (
	"fmt"
	"strings"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	EmailAddress string
}

func (e EmailNotifier) Send(message string) error {

	fmt.Printf("Sending email to %s\n", message)
	return nil
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) error {
	if !strings.HasPrefix(s.PhoneNumber, "06") {
		return fmt.Errorf("Invalid phone number")
	}

	fmt.Printf("Sending SMS to %s\n", message)
	return nil
}

type PushNotifier struct {
	DeviceID string
}

func (p PushNotifier) Send(message string) error {
	fmt.Printf("Sending Push Notification to %s\n", message)
	return nil
}

func main() {

	notifToSend := []Notifier{
		EmailNotifier{EmailAddress: "example@example.com"},
		SMSNotifier{PhoneNumber: "0612345678"},
		EmailNotifier{EmailAddress: "test@example.com"},
		SMSNotifier{PhoneNumber: "0712345678"},
		PushNotifier{DeviceID: "device123"},
	}

	message := "This is a test notification."

	for _, notifier := range notifToSend {
		err := notifier.Send(message)
		if err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
		}
	}
}
