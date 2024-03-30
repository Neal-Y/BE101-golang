package notifier

import (
	"be101_golang/models/user"
	"fmt"
)

type EmailNotifier struct{}

func (e EmailNotifier) Notify(user user.User, message string) {
	fmt.Println("Email sent to", user.GetName(), "with message:", message)
}

// func (e EmailNotifier) GetName() string {
// 	return "email"
// }
