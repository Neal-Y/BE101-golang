package user

import (
	"be101_golang/models/constant"
	"be101_golang/models/language"
)

type Guest struct {
	Name              string
	PreferredLanguage language.Language
}

func (g Guest) GetName() string                         { return constant.Guest }
func (g Guest) GetPreferredLanguage() language.Language { return g.PreferredLanguage }
