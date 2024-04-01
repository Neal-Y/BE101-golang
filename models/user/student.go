package user

import "be101_golang/models/language"

type Student struct {
	Name              string
	PreferredLanguage language.Language
}

func (s Student) GetName() string                         { return s.Name }
func (s Student) GetPreferredLanguage() language.Language { return s.PreferredLanguage }
