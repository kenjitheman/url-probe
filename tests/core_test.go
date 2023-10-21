package tests

import (
	"testing"
    "github.com/kenjitheman/url-probe/core"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "Valid URLs",
			args:      []string{"https://www.github.com", "https://www.google.com"},
			expectErr: false,
		},
		{
			name:      "No URLs Provided",
			args:      []string{},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := core.Run(test.args)
			if test.expectErr && err == nil {
				t.Errorf("expected an error but got none")
			} else if !test.expectErr && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}
