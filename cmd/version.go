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
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(), "cli-gopher v%s\n", resolveVersion())
		return nil
	},
}

func resolveVersion() string {
	if version == "" {
		return "dev"
	}
	return version
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
