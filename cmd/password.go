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
	Run:   passwordGenerator,
}

func passwordGenerator(cm *cobra.Command, args []string) {
	var r spg.CharRecipe
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

	for i := 0; i < passwordCount; i++ {
		p, err := r.Generate()
		if err != nil {
			fmt.Println(err)
			return
		}
		if passwordIsPin {
			fmt.Printf("Pin: %q\tEntropy: %.3f\n", p, p.Entropy)
		} else {
			fmt.Printf("Password: %q\tEntropy: %.3f\n", p, p.Entropy)
		}
	}
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	passwordCmd.Flags().IntVarP(&passwordCount, "count", "c", 1, "Number of passwords to generate")
	// TODO passwordCmd.Flags().IntVarP(&passwordCount, "exclude", "e", 1, "Exclude characters from generated passwords")
	passwordCmd.Flags().IntVarP(&passwordLength, "length", "l", 18, "Password length")
	passwordCmd.Flags().BoolVarP(&passwordIsPin, "pin", "p", false, "Create pin")
}
