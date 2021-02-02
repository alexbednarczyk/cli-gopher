package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version number of cli-gopher",
	Long:  `All software has versions. This is cli-gopher's`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-gopher v" + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
