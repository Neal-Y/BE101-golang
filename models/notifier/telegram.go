package notifier

import (
	"be101_golang/models/constant"
	"be101_golang/models/user"
	"fmt"
)

type TelegramNotifier struct{}

func (t TelegramNotifier) Notify(user user.User, message string) {
	fmt.Println("Telegram message sent to", user.GetName(), "with message:", message)
}

func (t TelegramNotifier) GetName() string {
	return constant.Telegram
}
