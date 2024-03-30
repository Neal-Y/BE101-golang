package language

import "be101_golang/models/constant"

type EnUS struct{}

func (e EnUS) GetLanguage() string {
	return constant.EnUSCode
}

func (e EnUS) GetMessage(event_name string) (string, error) {
	return LanguageCheck(e.GetLanguage(), event_name)
}
