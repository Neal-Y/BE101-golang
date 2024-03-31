package notifier

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
	"fmt"
)

type SMSNotifier struct{}

func (s SMSNotifier) Notify(user user.User, message string) {
	fmt.Println("SMS sent to", user.GetName(), "with message:", message)
}

func (s SMSNotifier) GetName() string {
	return constant.SMS
}
