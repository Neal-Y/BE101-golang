package language

import (
	"fmt"
)

type Language interface {
	GetMessage(event_name string) string
}

var EventMessage = map[string]map[string]string{
	"registerSuccess": {
		"zh-tw": "註冊成功。歡迎加入我們的社群！",
		"en-us": "Registration successful. Welcome to our community!",
	},
	"bookingSuccess": {
		"zh-tw": "課程預定成功。期待見到你！",
		"en-us": "Course successfully booked. We look forward to seeing you!",
	},
	"cancellationSuccess": {
		"zh-tw": "課程取消成功。希望在其他課程見到你。",
		"en-us": "Course successfully cancelled. We hope to see you in other courses.",
	},
}

type GenericLanguage struct {
	LanguageCode string
}

func (g GenericLanguage) GetMessage(event_name string) (string, error) {
	if _, ok := EventMessage[event_name][g.LanguageCode]; !ok {
		return "", fmt.Errorf("event name %s not found in %s", event_name, g.LanguageCode)
	}
	return EventMessage[event_name][g.LanguageCode], nil
}

/*
{
    "registerSuccess": {
        "en-us": "Registration successful. Welcome to our community!",
        "zh-tw": "註冊成功。歡迎加入我們的社群！"
    },
    "bookingSuccess": {
        "en-us": "Course successfully booked. We look forward to seeing you!",
        "zh-tw": "課程預定成功。期待見到你！"
     },
    "cancellationSuccess": {
        "en-us": "Course successfully cancelled. We hope to see you in other courses.",
        "zh-tw": "課程取消成功。希望在其他課程見到你。"
     },
}
*/
