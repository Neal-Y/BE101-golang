package user_test

import (
	"be101_golang/models/constant"
	"be101_golang/models/language"
	"be101_golang/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://go.dev/wiki/TableDrivenTests

func TestGetName(t *testing.T) {
	var flagtests = []struct {
		user   user.User
		expect string
	}{
		{user.Student{Name: "neal", PreferredLanguage: language.ZnTW{}}, "neal"},
		{user.Guest{}, "Guest"},
	}

	for _, tt := range flagtests {
		t.Run(tt.expect, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.user.GetName())
		})
	}
}

func TestGuestGetPreferredLanguage(t *testing.T) {
	var flagtests = []struct {
		user   user.User
		expect string // only string
	}{
		{user.Student{Name: "neal", PreferredLanguage: language.ZnTW{}}, constant.ZhTWCode},
		{user.Guest{PreferredLanguage: language.ZnTW{}}, constant.ZhTWCode},
	}

	for _, tt := range flagtests {
		t.Run(tt.expect, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.user.GetPreferredLanguage().GetLanguage())
		})
	}
}
