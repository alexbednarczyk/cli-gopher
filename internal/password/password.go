package password

import (
	"fmt"
	"go.1password.io/spg"
)

type GenerateOptions struct {
	Count  int
	Length int
	IsPin  bool
}

type GeneratedPassword struct {
	Password string
	Entropy  float32
	IsPin    bool
}

func Generate(opts GenerateOptions) ([]GeneratedPassword, error) {
	var r spg.CharRecipe
	if opts.IsPin {
		r = spg.CharRecipe{
			Length: opts.Length,
			Allow:  spg.Digits,
		}
	} else {
		r = spg.CharRecipe{
			Length:  opts.Length,
			Allow:   spg.Lowers | spg.Digits,
			Exclude: spg.Ambiguous,
		}
	}

	results := make([]GeneratedPassword, 0, opts.Count)
	for i := 0; i < opts.Count; i++ {
		p, err := r.Generate()
		if err != nil {
			return nil, err
		}
		results = append(results, GeneratedPassword{
			Password: p.String(),
			Entropy:  p.Entropy,
			IsPin:    opts.IsPin,
		})
	}
	return results, nil
}

func FormatOutput(passwords []GeneratedPassword) {
	for _, p := range passwords {
		if p.IsPin {
			fmt.Printf("Pin: %q\tEntropy: %.3f\n", p.Password, p.Entropy)
		} else {
			fmt.Printf("Password: %q\tEntropy: %.3f\n", p.Password, p.Entropy)
		}
	}
}
