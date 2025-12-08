package password

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	opts := GenerateOptions{
		Count:  5,
		Length: 10,
		IsPin:  false,
	}
	passwords, err := Generate(opts)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(passwords) != 5 {
		t.Errorf("Generate() returned %d passwords, want 5", len(passwords))
	}
	for _, p := range passwords {
		if len(p.Password) != 10 {
			t.Errorf("Generate() password length = %d, want 10", len(p.Password))
		}
	}
}
