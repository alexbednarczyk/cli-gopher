package environment

import (
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	type args struct {
		envIndex  int
		envPrefix string
	}
	tests := []struct {
		name    string
		args    args
		wantLog int8
	}{
		{
			name: "default",
			args: args{
				envIndex:  0,
				envPrefix: "",
			},
			wantLog: 1,
		},
		{
			name: "second environment",
			args: args{
				envIndex:  1,
				envPrefix: "",
			},
			wantLog: 2,
		},
		{
			name: "out of bounds (defaults to 0)",
			args: args{
				envIndex:  99,
				envPrefix: "",
			},
			wantLog: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnvironment(tt.args.envIndex, tt.args.envPrefix)
			if err != nil {
				t.Errorf("GetEnvironment() error = %v", err)
				return
			}
			if got.LogLevel != tt.wantLog {
				t.Errorf("GetEnvironment() LogLevel = %v, want %v", got.LogLevel, tt.wantLog)
			}
		})
	}
}
