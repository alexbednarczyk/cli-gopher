package count

import (
	"reflect"
	"testing"
)

func TestCountToInt(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  []int
	}{
		{
			name:  "zero",
			input: 0,
			want:  []int{0},
		},
		{
			name:  "positive integer",
			input: 3,
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "negative float",
			input: -2.5,
			want:  []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountToInt(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
