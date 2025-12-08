package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count [int to count to]",
	Short: "A simple count command",
	Long:  `A simple count command that counts from 0 to the input number. Prefix negatives with -- to avoid flag parsing (e.g. cli-gopher count -- -3).`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly 1 argument")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("argument must be integer")
		}
		return nil
	},
	RunE: countToInt,
}

func countToInt(cm *cobra.Command, args []string) error {
	countTo, _ := strconv.Atoi(args[0])
	if countTo < 0 {
		countTo = -countTo
	}

	writer := cm.OutOrStdout()
	fmt.Fprintln(writer)
	for i := 0; i <= countTo; i++ {
		fmt.Fprintf(writer, "%d", i)
		if i < countTo {
			fmt.Fprint(writer, " ")
		}
	}
	fmt.Fprintln(writer)
	return nil
}

func init() {
	rootCmd.AddCommand(countCmd)
}
