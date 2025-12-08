package cmd

import (
	"fmt"

	"github.com/alexbednarczyk/cli-gopher/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Simple config command",
	Long:  `Simple command that enables the user to read/write the cli-gopher config file`,
	Args:  cobra.NoArgs,
	Run:   runConfig,
}

func runConfig(cm *cobra.Command, args []string) {

	if len(args) == 0 {
		s, err := config.GetConfigString()
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Printf("\n%s\n", s)
		}
	}

}
func init() {
	rootCmd.AddCommand(configCmd)
}
