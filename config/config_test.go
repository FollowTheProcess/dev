package config_test

import (
	"strings"
	"testing"

	"github.com/FollowTheProcess/dev/config"
	"github.com/FollowTheProcess/test"
)

const (
	good = `
directory = "~/Development"

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
	in := strings.NewReader(good)
	cfg, err := config.Load(in)

	want := config.Config{
		GitHub: config.GitHub{
			Username: "FollowTheProcess",
			Token:    "notatoken",
		},
		Directory: "~/Development",
		Editor: config.Editor{
			Bin:  "code",
			Name: "VSCode",
			Open: true,
		},
	}

	test.Ok(t, err)
	test.Diff(t, cfg, want)
}
