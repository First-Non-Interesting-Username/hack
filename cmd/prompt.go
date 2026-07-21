package cmd

import (
	"fmt"
	"io"
	"os"
)

func GeneratePrompt() (string, error) {
	stdin, interactive, err := readPrompt()
	if err != nil {
		return "", err
	}
	switch {
	case prompt == "" && stdin == "":
		return "", fmt.Errorf("no prompt provided (use --prompt or stdin)")
	case prompt != "":
		return prompt, nil
	case stdin != "":
		return stdin, nil
	}

	promptText := prompt + "\n\n-----BEGIN STDIN CONTENT-----\n" + stdin + "\n-----END STDIN CONTENT-----"
	return promptText, nil
}

func readPrompt() (string, bool, error) {

	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", false, fmt.Errorf("checking stdin: %w", err)
	}

	isTerminal := (stat.Mode() & os.ModeCharDevice) != 0

	if isTerminal {
		fmt.Fprintln(os.Stderr, "Enter your prompt. Press Ctrl+D when done.")
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", false, fmt.Errorf("reading stdin: %w", err)
	}
	data = strings.TrimSpace(data)
	return string(data), isTerminal, nil
}
