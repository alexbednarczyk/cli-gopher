package cmd

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count [int to count to]",
	Short: "A simple count command",
	Long:  `A simple count command that counts from 0 to the input number`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires exactly 1 argument")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Argument must be integer")
		}
		return nil
	},
	Run: countToInt,
}

func countToInt(cm *cobra.Command, args []string) {
	floatCount, _ := strconv.ParseFloat(args[0], 64)
	countTo := int(math.Abs(floatCount))

	fmt.Printf("\n")
	for i := 0; i <= countTo; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
}

func init() {
	rootCmd.AddCommand(countCmd)
}
