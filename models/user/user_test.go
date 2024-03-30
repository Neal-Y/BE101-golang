package user_test

import (
	"be101_golang/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuestGetName(t *testing.T) {
	g := user.Guest{}

	if g.GetName() != "Guest" {
		t.Errorf("Expected Guest, got %s", g.GetName())
	}
	assert.Equal(t, "Guest", g.GetName())
}

// https://go.dev/wiki/TableDrivenTests

func TestGetName(t *testing.T) {
	var flagtests = []struct {
		user   user.User
		expect string
	}{
		{user.Student{}, "Student"},
		{user.Guest{}, "Guest"},
	}

	for _, tt := range flagtests {
		t.Run(tt.expect, func(t *testing.T) {

			if tt.user.GetName() != tt.expect {
				t.Errorf("Expected %s, got %s", tt.expect, tt.user.GetName())
			}
			assert.Equal(t, tt.expect, tt.user.GetName())

		})
	}
}

// func TestGuestGetPreferredLanguage(t *testing.T) {
// 	lang := user.ZnTW{}
// 	g := user.Guest{}
// 	if g.GetPreferredLanguage() != nil {
// 		t.Errorf("Expected nil, got %s", g.GetPreferredLanguage())
// 	}
// 	assert.Nil(t, g.GetPreferredLanguage())
// }
