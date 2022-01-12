package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidEnvVar(t *testing.T) {
	tests := []struct {
		description string
		input       string
		want        bool
	}{
		{
			description: "invalid: PATH",
			input:       "PATH=/opt/bin:$PATH",
			want:        false,
		},
		{
			description: "valid: HTTP_PROXY",
			input:       "HTTP_PROXY=http://localhost:8080",
			want:        true,
		},
		{
			description: "invalid: YGG_SOCKET_ADDR",
			input:       "YGG_SOCKET_ADDR=@yggd",
			want:        false,
		},
		{
			description: "valid: YGGD_VERSION",
			input:       "YGGD_VERSION=1.2.3",
			want:        true,
		},
		{
			description: "invalid: YGG_VERSION",
			input:       "YGG_VERSION=1.2.3",
			want:        false,
		},
		{
			description: "invalid: YGG_=",
			input:       "YGG_=",
			want:        false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := validEnvVar(test.input)

			if !cmp.Equal(got, test.want) {
				t.Errorf("%#v != %#v", got, test.want)
			}
		})
	}
}
