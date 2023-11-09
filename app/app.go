// Package app implements the functionality behind the CLI and allows us to execute it
// in tests without having to feed arguments in.
package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/FollowTheProcess/dev/config"
	"github.com/FollowTheProcess/msg"
)

// App represents the dev program.
type App struct {
	Stdout  io.Writer
	Stderr  io.Writer
	cfgPath string
	cfg     config.Config
	cfgOk   bool
}

// Config returns the set config.
func (a App) Config() config.Config {
	return a.cfg
}

// New builds and returns a new App.
//
// It should be called during building of every dev command and subcommand so each has
// access to the global state.
func New() (App, error) {
	fmt.Println("app.New called")
	app := App{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		cfgOk:  false,
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return App{}, fmt.Errorf("could not get user home dir: %w", err)
	}
	cfgFile := filepath.Join(home, ".dev.toml")
	app.cfgPath = cfgFile

	file, err := os.Open(cfgFile)
	if err != nil {
		// Here everything is fine apart from the config file which isn't a dealbreaker for
		// starting the app, so we allow it
		// TODO: Use a sensible default?
		msg.Fwarn(app.Stdout, "Config file %s missing or cannot be read: %v", cfgFile, err)
		return app, nil
	}
	defer file.Close()

	cfg, err := config.Load(file)
	if err != nil {
		// Here everything is fine apart from the config file which isn't a dealbreaker for
		// starting the app, so we allow it
		// TODO: Use a sensible default?
		msg.Fwarn(app.Stdout, "Config file %s cannot be read: %v", cfgFile, err)
		return app, nil
	}

	// By now we have loaded the config and everything is fine
	app.cfg = cfg
	app.cfgOk = true

	return app, nil
}
