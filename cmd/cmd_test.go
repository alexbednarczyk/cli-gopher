package cmd

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func executeCommand(t *testing.T, args ...string) (string, error) {
	t.Helper()

	viper.Reset()
	buf := new(bytes.Buffer)

	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetIn(bytes.NewBuffer(nil))
	rootCmd.SetArgs(args)

	err := rootCmd.Execute()

	rootCmd.SetArgs(nil)

	return buf.String(), err
}

func TestCountCommand(t *testing.T) {
	output, err := executeCommand(t, "count", "3")
	require.NoError(t, err)
	require.Equal(t, "\n0 1 2 3\n", output)
}

func TestCountCommandHandlesNegative(t *testing.T) {
	output, err := executeCommand(t, "count", "--", "-2")
	require.NoError(t, err)
	require.Equal(t, "\n0 1 2\n", output)
}

func TestDateCommand(t *testing.T) {
	output, err := executeCommand(t, "date")
	require.NoError(t, err)

	dateStr := strings.TrimSpace(output)
	_, parseErr := time.Parse("2006-01-02", dateStr)
	require.NoError(t, parseErr)
}

func TestEnvCommandDefaults(t *testing.T) {
	output, err := executeCommand(t, "env")
	require.NoError(t, err)

	require.Contains(t, output, "Debug:true")
	require.Contains(t, output, "LogLevel:1")
	require.Contains(t, output, "Rate:1.2")
}

func TestEnvCommandWithPrefix(t *testing.T) {
	t.Setenv("DEMO_DEBUG", "false")
	t.Setenv("DEMO_LOG_LEVEL", "3")
	t.Setenv("DEMO_RATE", "4.2")

	output, err := executeCommand(t, "env", "--prefix", "DEMO", "--environment", "1")
	require.NoError(t, err)

	require.Contains(t, output, "Debug:false")
	require.Contains(t, output, "LogLevel:3")
	require.Contains(t, output, "Rate:4.2")
}

func TestPasswordCommandPin(t *testing.T) {
	output, err := executeCommand(t, "password", "--count", "2", "--length", "6", "--pin")
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(output), "\n")
	require.Len(t, lines, 2)

	for _, line := range lines {
		require.True(t, strings.HasPrefix(line, "Pin:"))
		require.Contains(t, line, "Entropy:")
	}
}

func TestVersionCommand(t *testing.T) {
	previous := version
	version = "1.2.3"
	t.Cleanup(func() {
		version = previous
	})

	output, err := executeCommand(t, "version")
	require.NoError(t, err)
	require.Equal(t, "cli-gopher v1.2.3\n", output)
}
