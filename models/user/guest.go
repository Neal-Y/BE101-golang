package user

import "be101_golang/models/language"

type Guest struct {
	PreferredLanguage language.Language
}

func (g Guest) GetName() string                         { return "Guest" }
func (g Guest) GetPreferredLanguage() language.Language { return g.PreferredLanguage }
