package event_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO : find that how to see the test coverage of the package

func TestEvent(t *testing.T) {
	var flagtests = []struct {
		in  string
		out string
	}{
		{"neal", "neal"},
		{"Guest", "Guest"},
	}

	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in)
		})
	}
}
