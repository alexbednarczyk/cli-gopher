package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Simple config command",
	Long:  `Simple command that enables the user to read/write the cli-gopher config file`,
	Args:  cobra.NoArgs,
	RunE:  config,
}

func config(cm *cobra.Command, args []string) error {
	completeConfig := viper.AllSettings()

	b, err := json.MarshalIndent(completeConfig, "", " ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	fmt.Fprintf(cm.OutOrStdout(), "\n%s\n", string(b))
	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)
}
