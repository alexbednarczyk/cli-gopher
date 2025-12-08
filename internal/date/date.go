package date

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Executor interface allows mocking the command execution
type Executor interface {
	Execute(name string, arg ...string) error
}

type RealExecutor struct{}

func (e *RealExecutor) Execute(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// GetDateCommand returns the command to execute based on the OS.
func GetDateCommand(osName string) (string, []string, error) {
	switch osName {
	case "darwin", "linux":
		return "date", nil, nil
	case "windows":
		return "date", []string{"/t"}, nil
	default:
		return "", nil, fmt.Errorf("platform is unsupported! Can't display date for %s", osName)
	}
}

// DisplayDate executes the date command for the current OS.
func DisplayDate(executor Executor) error {
	name, args, err := GetDateCommand(runtime.GOOS)
	if err != nil {
		return err
	}
	return executor.Execute(name, args...)
}
