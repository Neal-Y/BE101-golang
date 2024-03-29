package main

import (
	"be101_golang/models/event"
	"be101_golang/models/language"
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

func main() {
	// 產生notifier
	tel := notifier.TelegramNotifier{}
	sms := notifier.SMSNotifier{}
	email := notifier.EmailNotifier{}

	// 先產生人物

	// 產生一個 Guest
	guest := user.Guest{Name: "Guest", PreferredLanguage: language.ZnTW{}}
	// 產生一個 Student
	student := user.Student{Name: "neal", PreferredLanguage: language.EnUS{}}

	// 產生容器好了
	// which notifier want to notify (maybe SMS, Telegram, Email or all)
	//TODO: 限制容器只能裝這三個notifier，且不能重複
	notifier := &event.EventFactory{}
	notifier.AddNotifier(tel)
	notifier.AddNotifier(sms)
	notifier.AddNotifier(email)

	// 發生啥事
	signup := event.Signup{Methods: notifier}
	bookclass := event.BookClass{Methods: notifier}

	// 通知
	signup.Trigger(guest)
	bookclass.Trigger(student)
}
