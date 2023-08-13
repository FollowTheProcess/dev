package main

import (
	"os"

	"github.com/FollowTheProcess/dev/cli"
	"github.com/FollowTheProcess/msg"
)

func main() {
	if err := cli.Build().Execute(); err != nil {
		msg.Error("%s", err)
		os.Exit(1)
	}
}
