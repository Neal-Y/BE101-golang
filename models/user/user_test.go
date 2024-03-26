package user_test

import (
	"be101_golang/models/user"
	"testing"
)

func TestGuestGetName(t *testing.T) {
	g := user.Guest{}
	if g.GetName() != "Guest" {
		t.Errorf("Expected Guest, got %s", g.GetName())
	}
}
