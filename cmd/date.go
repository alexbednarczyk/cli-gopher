package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A simple date command",
	Long:  `A simple date command that displays today's date`,
	Args:  cobra.NoArgs,
	RunE:  todaysDate,
}

var dateFormat string

func todaysDate(cm *cobra.Command, args []string) error {
	now := time.Now().Format(dateFormat)
	fmt.Fprintln(cm.OutOrStdout(), now)
	return nil
}

func init() {
	rootCmd.AddCommand(dateCmd)
	dateCmd.Flags().StringVar(&dateFormat, "format", "2006-01-02", "Go time format for date output")
}
