package cmd

import (
	"encoding/json"
	"fmt"
	"log"
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
	Short: "A simple environvment command",
	Long:  `A simple environvment command that displays environment variables`,
	Args:  cobra.NoArgs,
	Run:   getAndDisplayEnv,
}

func getAndDisplayEnv(cm *cobra.Command, args []string) {

	var e environment

	if len(specs)-1 >= envEnvironment {
		e = specs[envEnvironment]
	} else {
		fmt.Println("Environment not found, using default")
		e = specs[0]
	}

	err := envconfig.Process(envPrefix, &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	s, _ := json.Marshal(e)
	unJSONify := strings.NewReplacer("{", "", "}", "", ",", "\n", "\"", "")
	output := unJSONify.Replace(string(s))
	fmt.Printf("------------\nKey:Value\n------------\n")
	fmt.Println(output)
}

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().StringVarP(&envPrefix, "prefix", "p", "", "Will use prefixed environment variable if it is set")
	envCmd.Flags().IntVarP(&envEnvironment, "environment", "e", 0, "Environment for environment variables")

	specs = append(specs, environment{Debug: true, LogLevel: 1, Rate: 1.2})
	specs = append(specs, environment{Debug: false, LogLevel: 2, Rate: 2.3})
}
