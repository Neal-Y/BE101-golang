package notifier

import (
	"be101_golang/models/user"
	"fmt"
)

type TelegramNotifier struct{}

func (t TelegramNotifier) Notify(user user.User, message string) {
	fmt.Println("Telegram message sent to", user.GetName(), "with message:", message)
}
