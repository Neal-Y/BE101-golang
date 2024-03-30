package language

import (
	"be101_golang/models/constant"
	"fmt"
)

type Language interface {
	GetMessage(event_name string) (string, error)
	GetLanguage() string
}

var EventMessage = map[string]map[string]string{
	constant.Newyear: {
		constant.ZhTWCode: "新年快樂",
		constant.EnUSCode: "Happy New Year",
	},
	constant.Register: {
		constant.ZhTWCode: "註冊成功。歡迎加入我們的社群！",
		constant.EnUSCode: "Registration successful. Welcome to our community!",
	},
	constant.Booking: {
		constant.ZhTWCode: "課程預定成功。期待見到你！",
		constant.EnUSCode: "Course successfully booked. We look forward to seeing you!",
	},
	constant.Cancel: {
		constant.ZhTWCode: "課程取消成功。希望在其他課程見到你。",
		constant.EnUSCode: "Course successfully cancelled. We hope to see you in other courses.",
	},
}

func LanguageCheck(lang string, event_name string) (message string, err error) {
	langMessage, ok := EventMessage[event_name]
	if !ok {
		return "", fmt.Errorf("event %s not found", event_name)
	}

	message, ok = langMessage[lang]
	if !ok {
		return "", fmt.Errorf("language %s not found", lang)
	}

	return
}
