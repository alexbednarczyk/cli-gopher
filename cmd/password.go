package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.1password.io/spg"
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
	RunE:  passwordGenerator,
}

func passwordGenerator(cm *cobra.Command, args []string) error {
	var r spg.CharRecipe

	if passwordLength <= 0 {
		return fmt.Errorf("length must be positive")
	}
	if passwordCount <= 0 {
		return fmt.Errorf("count must be positive")
	}

	if passwordIsPin {
		r = spg.CharRecipe{
			Length: passwordLength,
			Allow:  spg.Digits,
		}
	} else {
		r = spg.CharRecipe{
			Length:  passwordLength,
			Allow:   spg.Lowers | spg.Digits,
			Exclude: spg.Ambiguous,
		}
	}

	writer := cm.OutOrStdout()
	for i := 0; i < passwordCount; i++ {
		p, err := r.Generate()
		if err != nil {
			return fmt.Errorf("generate password: %w", err)
		}
		if passwordIsPin {
			fmt.Fprintf(writer, "Pin: %q\tEntropy: %.3f\n", p, p.Entropy)
		} else {
			fmt.Fprintf(writer, "Password: %q\tEntropy: %.3f\n", p, p.Entropy)
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	passwordCmd.Flags().IntVarP(&passwordCount, "count", "c", 1, "Number of passwords to generate")
	passwordCmd.Flags().IntVarP(&passwordLength, "length", "l", 18, "Password length")
	passwordCmd.Flags().BoolVarP(&passwordIsPin, "pin", "p", false, "Create pin")
}
