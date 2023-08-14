// Package config implements dev's configuration management.
package config

import (
	"errors"
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
)

// Config encodes dev's currently loaded configuration.
type Config struct {
	// Github config
	GitHub GitHub `toml:"github,omitempty"`
	// The absolute path to where user projects are stored
	Directory string `toml:"directory,omitempty"`
	// Editor config
	Editor Editor `toml:"editor,omitempty"`
}

// Editor encodes config relating to opening files/projects in the user's editor.
type Editor struct {
	// The name of the binary to use e.g. `code`
	Bin string `toml:"bin,omitempty"`
	// Friendly name of the editor to use for printed messages
	Name string `toml:"name,omitempty"`
	// Whether or not we should attempt to open things at all
	Open bool `toml:"open,omitempty"`
}

// GitHub encodes config relating to the user's GitHub credentials/identity.
type GitHub struct {
	// The user's GitHub user name
	Username string `toml:"username,omitempty"`
	// A personal access token with at least repo scope
	Token string `toml:"token,omitempty"`
}

// Load reads toml config from the reader and returns a Config.
func Load(r io.Reader) (Config, error) {
	var cfg Config
	if _, err := toml.NewDecoder(r).Decode(&cfg); err != nil {
		var parseError toml.ParseError
		if errors.As(err, &parseError) {
			return Config{}, fmt.Errorf("invalid config file: %s", parseError.ErrorWithPosition())
		}
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}

	return cfg, nil
}
