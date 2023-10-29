package config_test

import (
	"strings"
	"testing"

	"github.com/FollowTheProcess/dev/config"
	"github.com/FollowTheProcess/test"
)

const (
	good = `
directory = "/Users/me/Development"

[github]
username = "FollowTheProcess"
token = "notatoken"

[editor]
open = true
bin = "code"
name = "VSCode"
`
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    config.Config
		wantErr bool
	}{
		{
			name:  "happy path",
			input: good,
			want: config.Config{
				GitHub: config.GitHub{
					Username: "FollowTheProcess",
					Token:    "notatoken",
				},
				Directory: "/Users/me/Development",
				Editor: config.Editor{
					Bin:  "code",
					Name: "VSCode",
					Open: true,
				},
			},
			wantErr: false,
		},
		{
			name:    "minimal",
			input:   `directory = "/Users/me/somewhere/else"`,
			want:    config.Config{Directory: "/Users/me/somewhere/else"},
			wantErr: false,
		},
		{
			name:    "empty",
			input:   "", // Only default should be Directory
			want:    config.Config{},
			wantErr: false,
		},
		{
			name:    "syntax error",
			input:   "some nonsense [(*&^)]",
			want:    config.Config{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := strings.NewReader(tt.input)
			got, err := config.Load(in)

			test.WantErr(t, err, tt.wantErr)
			test.Diff(t, got, tt.want)
		})
	}
}
