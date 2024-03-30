package language

import "be101_golang/models/constant"

type ZnTW struct{}

func (z ZnTW) GetLanguage() string {
	return constant.ZhTWCode
}

func (z ZnTW) GetMessage(event_name string) (string, error) {
	return LanguageCheck(z.GetLanguage(), event_name)
}
