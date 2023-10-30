// Package app implements the functionality behind the CLI and allows us to execute it
// in tests without having to feed arguments in.
package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/FollowTheProcess/dev/config"
	"github.com/FollowTheProcess/msg"
)

// onceNew synchronises the call to app.New, ensuring we only ever set it up once.
var onceNew = sync.OnceValues(newApp)

// App represents the dev program.
type App struct {
	Stdout io.Writer
	Stderr io.Writer
	cfg    config.Config
	cfgOk  bool
}

// Config returns the set config.
func (a App) Config() config.Config {
	return a.cfg
}

// New builds and returns a new App.
//
// It is safe to call multiple times and concurrently as under the hood
// it is synchronised by sync.Once, ensuring that the app is only ever constructed
// once, including reading and parsing the config file.
//
// It should be called during building of every dev command and subcommand so each has
// access to the global state.
func New() (App, error) {
	return onceNew()
}

// newApp builds and returns a new app.
func newApp() (App, error) {
	fmt.Println("app.New() called") // So I can feel warm and fuzzy that it only runs once
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
