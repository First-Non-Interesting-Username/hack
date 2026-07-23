package cmd

import (
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockStdin(content string) func() {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	w.Write([]byte(content))
	w.Close()
	os.Stdin = r
	return func() { os.Stdin = oldStdin }
}

func TestGenerateSystemPrompt(t *testing.T) {
	tests := []struct {
		name        string
		command     bool
		code        bool
		expectWords []string
	}{
		{"Normal mode", false, false, []string{"terminal environment"}},
		{"Command mode", true, false, []string{"COMMAND MODE", "UNIX filter"}},
		{"Code mode", false, true, []string{"CODE MODE", "source code"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandMode = tt.command
			codeMode = tt.code

			systemPrompt, err := generateSystemPrompt()
			require.NoError(t, err)

			for _, word := range tt.expectWords {
				assert.True(t, strings.Contains(systemPrompt, word), "Prompt should contain: "+word)
			}
		})
	}
}

func TestGeneratePrompt(t *testing.T) {
	restore := mockStdin("some piped data")
	defer restore()

	prompt = "Summarize this"

	result, err := generatePrompt()
	require.NoError(t, err)

	assert.Contains(t, result, "Summarize this", "Should contain the explicit prompt")
	assert.Contains(t, result, "-----BEGIN STDIN CONTENT-----", "Should contain stdin delimiters")
	assert.Contains(t, result, "some piped data", "Should contain the piped data")
}
