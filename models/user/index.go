package user

import "be101_golang/models/language"

// User is the method set
type User interface {
	GetName() string
	GetPreferredLanguage() language.Language
}
