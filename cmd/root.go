package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version     = "dev"
	cfgFile     string
	cfgFileName = ".gopher"
)

var rootCmd = &cobra.Command{
	Use:   "cli-gopher",
	Short: "Example cli tool based on golang",
	Long:  `Example cli tool based on golang!`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig()
	},
}

// Execute function
func Execute() error {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .gopher.yaml in current directory)")
}

func initConfig() error {
	viper.SetEnvPrefix("GOPHER")
	viper.AutomaticEnv()
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName(cfgFileName)
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var cfgErr viper.ConfigFileNotFoundError
		if errors.As(err, &cfgErr) {
			return nil
		}
		return fmt.Errorf("read config: %w", err)
	}

	return nil
}
