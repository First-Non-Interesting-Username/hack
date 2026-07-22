package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "1",
	Use:     "hack",
	Short:   "CLI tool for interacting with LLMs",
	Long: `
hack is a command-line tool for AI-assisted development.

It sends prompts to OpenAI-compatible LLM APIs and returns output tailored
to the task at hand: shell commands, executable code, documentation, or
plain responses.

Usage:
	hack -p "your prompt here"
	echo "some content" | hack -p "summarize this"
	ls -la | hack -x "delete the largest file"

Modes:
	shell	(-x/--execute)	Generate shell commands from a prompt
	code	(-w/--write)  	Output executable code (jq, python3, bash, or POSIX sh)
	normal					Standard prompt-and-response
	`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cfgPath, _ := cmd.Flags().GetString("config")
		if err := createConfig(cfgPath); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runHack()
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runHack() error {
	response, err := makeRequest()
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
