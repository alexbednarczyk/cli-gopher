package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version     string
	cfgFile     string
	cfgFileName = ".gopher"
)

var rootCmd = &cobra.Command{
	Use:   "cli-gopher",
	Short: "Example cli tool based on golang",
	Long:  `Example cli tool based on golang!`,
}

// Execute function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".gopher" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgFileName)
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				initConfigPath := home + string(os.PathSeparator) + cfgFileName + ".yaml"
				if err := ioutil.WriteFile(initConfigPath, []byte(""), 0666); err != nil {
					fmt.Println("error:", err)
				}
				fmt.Printf("Config file initialized: %s\n", initConfigPath)
			} else {
				// Config file was found but another error was produced
				er(err)
			}

		}

	}

	// Set defaults for config file
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	if err := viper.WriteConfig(); err != nil {
		fmt.Println("error:", err)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
