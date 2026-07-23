package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetermineMode(t *testing.T) {
	tests := []struct {
		name        string
		commandFlag bool
		codeFlag    bool
		expected    string
		expectError bool
	}{
		{"Normal mode when no flags are set", false, false, "normal", false},
		{"Command mode when shell flag is set", true, false, "command", false},
		{"Code mode when write flag is set", false, true, "code", false},
		{"Error when both flags are set", true, true, "", true},
	}

	for _, test := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mode, err := determineMode(test.commandFlag, test.codeFlag)

			if tt.expectError {
                assert.Error(t, err, "Expected an error for conflicting flags")
            } else {
                assert.NoError(t, err, "Did not expect an error")
                assert.Equal(t, tt.expected, mode, "Mode did not match expected output")
            }
		})
	}
}
