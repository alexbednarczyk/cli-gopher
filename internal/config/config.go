package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

func GetConfigString() (string, error) {
	completeConfig := viper.AllSettings()

	b, err := json.MarshalIndent(completeConfig, "", " ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
