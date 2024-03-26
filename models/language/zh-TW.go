package language

type ZnTW struct{}

func (z ZnTW) GetLanguage() string {
	return ZhTWCode
}

func (z ZnTW) GetMessage(event_name string) (string, error) {
	return LanguageCheck(z.GetLanguage(), event_name)
}
