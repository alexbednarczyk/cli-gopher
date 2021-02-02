package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var theDate map[string]func()

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A simple date command",
	Long:  `A simple date command that displays today's date`,
	Args:  cobra.NoArgs,
	Run:   todaysDate,
}

func todaysDate(cm *cobra.Command, args []string) {
	value, ok := theDate[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Platform is unsupported! Can't display date for " + runtime.GOOS)
	}
}

func init() {
	rootCmd.AddCommand(dateCmd)

	theDate = make(map[string]func())
	theDate["darwin"] = func() {
		cmd := exec.Command("date")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	theDate["linux"] = func() {
		cmd := exec.Command("date")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	theDate["windows"] = func() {
		cmd := exec.Command("date", "/t")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
