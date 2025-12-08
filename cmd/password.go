package cmd

import (
	"fmt"

	"github.com/alexbednarczyk/cli-gopher/internal/password"
	"github.com/spf13/cobra"
)

var (
	passwordCount  int
	passwordLength int
	passwordIsPin  bool
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "A simple password command",
	Long:  `A simple password command that uses 1Password's spg golang library. http://github.com/1Password/spg`,
	Run:   passwordGenerator,
}

func passwordGenerator(cm *cobra.Command, args []string) {
	opts := password.GenerateOptions{
		Count:  passwordCount,
		Length: passwordLength,
		IsPin:  passwordIsPin,
	}

	passwords, err := password.Generate(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	password.FormatOutput(passwords)
}
func init() {
	rootCmd.AddCommand(passwordCmd)

	passwordCmd.Flags().IntVarP(&passwordCount, "count", "c", 1, "Number of passwords to generate")
	passwordCmd.Flags().IntVarP(&passwordLength, "length", "l", 18, "Password length")
	passwordCmd.Flags().BoolVarP(&passwordIsPin, "pin", "p", false, "Create pin")
}
