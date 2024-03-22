package user

import "be101_golang/models/language"

// User is the method set
type User interface {
	GetName() string
	GetPreferredLanguage() string
}

type Student struct {
	Name              string
	PreferredLanguage language.GenericLanguage
}
type Guest struct {
	Name              string
	PreferredLanguage language.GenericLanguage
}

func (s Student) GetName() string              { return s.Name }
func (s Student) GetPreferredLanguage() string { return s.PreferredLanguage.LanguageCode }

func (g Guest) GetName() string              { return "Guest" }
func (g Guest) GetPreferredLanguage() string { return g.PreferredLanguage.LanguageCode }
