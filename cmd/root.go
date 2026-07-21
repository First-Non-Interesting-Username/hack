package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "1",
	Use:     "hack",
	Short:   "CLI tool for interacting with LLMs",
	Long: `hack is a CLI tool for programmatically interacting with Open AI compatible APIs, with main focus on Hack Club AI.

	It features 3 modes: normal, command and code.

	TBD
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
