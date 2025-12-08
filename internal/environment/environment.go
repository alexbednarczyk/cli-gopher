package environment

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Debug    bool    `envconfig:"DEBUG"`
	LogLevel int8    `envconfig:"LOG_LEVEL"`
	Rate     float64 `envconfig:"RATE"`
}

var defaultSpecs = []Environment{
	{Debug: true, LogLevel: 1, Rate: 1.2},
	{Debug: false, LogLevel: 2, Rate: 2.3},
}

// GetEnvironment returns the environment configuration based on the index and prefix.
func GetEnvironment(envIndex int, envPrefix string) (Environment, error) {
	var e Environment

	if len(defaultSpecs)-1 >= envIndex {
		e = defaultSpecs[envIndex]
	} else {
		// Log warning? For now just use default
		e = defaultSpecs[0]
	}

	err := envconfig.Process(envPrefix, &e)
	if err != nil {
		return Environment{}, err
	}
	return e, nil
}

// FormatEnvironment formats the environment struct into the desired string format.
func FormatEnvironment(e Environment) (string, error) {
	s, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	unJSONify := strings.NewReplacer("{", "", "}", "", ",", "\n", "\"", "")
	output := unJSONify.Replace(string(s))
	return fmt.Sprintf("------------\nKey:Value\n------------\n%s", output), nil
}
