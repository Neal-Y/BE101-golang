package language

import "fmt"

type Language interface {
	GetMessage(event_name string) (string, error)
}

//TODO: use this list to validate the language input
// var LanguageList = []string{"zh-tw", "en-us"}

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

type ZnTW struct{}
type EnUS struct{}

// TODO: improve this function with a switch statement, use other flow control statements.
func NewLanguage(lang string) (Language, error) {
	switch lang {
	case "zh-tw":
		return ZnTW{}, nil
	case "en-us":
		return EnUS{}, nil
	default:
		return nil, fmt.Errorf("language %s not found", lang)
	}
}

func (z ZnTW) GetMessage(event_name string) (string, error) {
	return LanguageCheck("zh-tw", event_name)
}

func (e EnUS) GetMessage(event_name string) (string, error) {
	return LanguageCheck("en-us", event_name)
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
