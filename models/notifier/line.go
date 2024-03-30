package notifier

import (
	"be101_golang/models/user"
	"fmt"
)

type LineNotifier struct{}

func (l LineNotifier) Notify(user user.User, message string) {
	fmt.Println("Line sent to", user.GetName(), "with message:", message)
}
