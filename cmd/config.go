package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	commandMode bool
	codeMode    bool
	Prompt      string
	showVersion bool
)

func init() {
	rootCmd.Flags().StringVarP(&Prompt, "prompt", "p", "", "prompt for the LLM")
	rootCmd.Flags().StringP("config", "c", "", "config file path, $HOME/.config/hack-ai/config.toml if not provided")
	rootCmd.Flags().StringP("key", "k", "", "API key for selected provider")
	rootCmd.Flags().StringP("model", "m", "", "LLM used for response")
	rootCmd.Flags().StringP("base", "b", "", "base URL for the API (without /chat/completions)")
	rootCmd.Flags().BoolVarP(&commandMode, "shell", "s", false, "enable command mode")
	rootCmd.Flags().BoolVarP(&codeMode, "write", "w", false, "enable code mode")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "display informations about version")

	viper.BindPFlag("api_key", rootCmd.Flags().Lookup("key"))
	viper.BindPFlag("model", rootCmd.Flags().Lookup("model"))
	viper.BindPFlag("base_url", rootCmd.Flags().Lookup("base"))
}

func createConfig(cmd *cobra.Command) error {
	cfgPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
	} else {
		configDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		viper.AddConfigPath(filepath.Join(configDir, "hack-ai"))
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}
}

func determineMode() (string, error) {
	if commandMode && codeMode {
		return "", fmt.Errorf("cannot use command mode and code mode simultaneously")
	}

	switch {
	case commandMode:
		return "command", nil
	case codeMode:
		return "code", nil
	default:
		return "normal", nil
	}
}

func verifyflags(cmd *cobra.Command) error {
	if Prompt == "" && showVersion == false && cmd.Flags().GetBool("help") == false {
		return fmt.Errorf("no prompt was provided")
	}
}
