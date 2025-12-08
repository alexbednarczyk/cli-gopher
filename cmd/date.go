package cmd

import (
	"fmt"
	"github.com/alexbednarczyk/cli-gopher/internal/date"
	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A simple date command",
	Long:  `A simple date command that displays today's date`,
	Args:  cobra.NoArgs,
	Run:   todaysDate,
}

func todaysDate(cm *cobra.Command, args []string) {
	executor := &date.RealExecutor{}
	if err := date.DisplayDate(executor); err != nil {
		fmt.Println("error:", err)
	}
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
