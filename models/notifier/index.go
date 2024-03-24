package notifier

import (
	"be101_golang/models/user"
	"fmt"
)

type Notifier interface {
	Notify(user user.User, message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}
type TelegramNotifier struct{}

func (e EmailNotifier) Notify(user user.User, message string) {
	fmt.Println("Email sent to", user.GetName(), "with message:", message)
}
func (s SMSNotifier) Notify(user user.User, message string) {
	fmt.Println("SMS sent to", user.GetName(), "with message:", message)
}
func (t TelegramNotifier) Notify(user user.User, message string) {
	fmt.Println("Telegram message sent to", user.GetName(), "with message:", message)
}
