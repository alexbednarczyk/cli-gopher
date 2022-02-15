package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version     string
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
	// Search config in present working directory with name ".gopher" (without extension).
	viper.AddConfigPath("./")
	viper.SetConfigName(cfgFileName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			initConfigPath := "./" + string(os.PathSeparator) + cfgFileName + ".yaml"
			if err := ioutil.WriteFile(initConfigPath, []byte(""), 0666); err != nil {
				fmt.Println("error:", err)
			}
			// Set defaults for config file
			viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
			viper.SetDefault("license", "apache")
			fmt.Printf("Config file initialized: %s\n", initConfigPath)
		} else {
			// Config file was found but another error was produced
			er(err)
		}
	}

	if err := viper.WriteConfig(); err != nil {
		fmt.Println("error:", err)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
