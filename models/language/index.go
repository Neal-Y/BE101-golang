package language

import "fmt"

type Language interface {
	GetMessage(event_name string) (string, error)
	GetLanguage() string
}

const (
	EnUSCode = "en-us"
	ZhTWCode = "zh-tw"
)

var EventMessage = map[string]map[string]string{
	"register": {
		ZhTWCode: "註冊成功。歡迎加入我們的社群！",
		EnUSCode: "Registration successful. Welcome to our community!",
	},
	"booking": {
		ZhTWCode: "課程預定成功。期待見到你！",
		EnUSCode: "Course successfully booked. We look forward to seeing you!",
	},
	"cancellation": {
		ZhTWCode: "課程取消成功。希望在其他課程見到你。",
		EnUSCode: "Course successfully cancelled. We hope to see you in other courses.",
	},
}

func LanguageCheck(lang string, event_name string) (string, error) {
	if langMessage, ok := EventMessage[event_name]; ok {
		if message, ok := langMessage[lang]; ok {
			return message, nil
		}
	}
	return "", fmt.Errorf("message %s not found", event_name)
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
