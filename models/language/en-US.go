package language

type EnUS struct{}

func (e EnUS) GetLanguage() string {
	return EnUSCode
}

func (e EnUS) GetMessage(event_name string) (string, error) {
	return LanguageCheck(e.GetLanguage(), event_name)
}
