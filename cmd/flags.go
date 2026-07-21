package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var (
	Prompt string
	ConfigPath string
	ApiKey string
	Model string
	BaseURL string
	CommandMode bool
	CodeMode bool
)

func init() {
	rootCmd.Flags().StringVarP(&Prompt, "prompt", "p", "", "prompt for the LLM"),
	rootCmd.Flags().StringVarP(&ConfigPath, "config", "c", "","config file path, $HOME/.config/hack-ai/config.toml if not provided"),
	rootCmd.Flags().StringVarP(&ApiKey, "key", "k", "", "API key for selected provider"),
	rootCmd.Flags().StringVarP(&Model, "model", "m", "", "LLM used for response"),
	rootCmd.Flags().StringVarP(&BaseURL, "base", "b", "", "base URL for the API (without /chat/completions)"),
	rootCmd.Flags().BoolVarP(&CommandMode, "shell", "s", "", "enable command mode"),
	rootCmd.Flags().BoolVarP(&CodeMode, "shell", "s", "", "enable code mode"),
}
