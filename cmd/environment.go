package cmd

import (
	"fmt"
	"log"

	"github.com/alexbednarczyk/cli-gopher/internal/environment"
	"github.com/spf13/cobra"
)

var (
	envPrefix      string
	envEnvironment int
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A simple environvment command",
	Long:  `A simple environvment command that displays environment variables`,
	Args:  cobra.NoArgs,
	Run:   getAndDisplayEnv,
}

func getAndDisplayEnv(cm *cobra.Command, args []string) {
	e, err := environment.GetEnvironment(envEnvironment, envPrefix)
	if err != nil {
		log.Fatal(err.Error())
	}

	output, err := environment.FormatEnvironment(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().StringVarP(&envPrefix, "prefix", "p", "", "Will use prefixed environment variable if it is set")
	envCmd.Flags().IntVarP(&envEnvironment, "environment", "e", 0, "Environment for environment variables")
}
