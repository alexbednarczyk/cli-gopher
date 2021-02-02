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
	Run:   config,
}

func config(cm *cobra.Command, args []string) {

	if len(args) == 0 {
		completeConfig := viper.AllSettings()

		if b, err := json.MarshalIndent(completeConfig, "", " "); err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("\n%s\n", string(b))
		}
	}

}

func init() {
	rootCmd.AddCommand(configCmd)
}
