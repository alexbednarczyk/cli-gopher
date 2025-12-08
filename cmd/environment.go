package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
)

var (
	envPrefix      string
	envEnvironment int
	specs          []environment
)

type environment struct {
	Debug    bool    `envconfig:"DEBUG"`
	LogLevel int8    `envconfig:"LOG_LEVEL"`
	Rate     float64 `envconfig:"RATE"`
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A simple environment command",
	Long:  `A simple environment command that displays environment variables`,
	Args:  cobra.NoArgs,
	RunE:  getAndDisplayEnv,
}

func getAndDisplayEnv(cm *cobra.Command, args []string) error {
	var e environment

	if envEnvironment >= 0 && len(specs)-1 >= envEnvironment {
		e = specs[envEnvironment]
	} else {
		fmt.Fprintln(cm.ErrOrStderr(), "Environment not found, using default")
		e = specs[0]
	}

	err := envconfig.Process(envPrefix, &e)
	if err != nil {
		return fmt.Errorf("read environment variables: %w", err)
	}

	s, _ := json.Marshal(e)
	unJSONify := strings.NewReplacer("{", "", "}", "", ",", "\n", "\"", "")
	output := unJSONify.Replace(string(s))
	writer := cm.OutOrStdout()
	fmt.Fprintln(writer, "------------")
	fmt.Fprintln(writer, "Key:Value")
	fmt.Fprintln(writer, "------------")
	fmt.Fprintln(writer, output)
	return nil
}

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().StringVarP(&envPrefix, "prefix", "p", "", "Will use prefixed environment variable if it is set")
	envCmd.Flags().IntVarP(&envEnvironment, "environment", "e", 0, "Environment for environment variables")

	specs = append(specs, environment{Debug: true, LogLevel: 1, Rate: 1.2})
	specs = append(specs, environment{Debug: false, LogLevel: 2, Rate: 2.3})
}
