package main

import (
	"os"
	"runtime/debug"

	"github.com/FollowTheProcess/dev/cli"
	"github.com/FollowTheProcess/msg"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			msg.Error("dev panicked, this is a bug!\nCausing error:\t%v\nStack:\n%s", r, debug.Stack())
			os.Exit(1)
		}
	}()
	if err := cli.Build().Execute(); err != nil {
		msg.Error("%s", err)
		os.Exit(1) //nolint:gocritic // Exit only called in err case, deferred func is to catch panics (exitAfterDefer)
	}
}
