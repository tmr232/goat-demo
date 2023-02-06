package main

import (
	"github.com/tmr232/goat"
	"github.com/tmr232/goat/flags"
	"github.com/urfave/cli/v2"
)

func init() {
	goat.Register(Hello, goat.RunConfig{
		Flags: []cli.Flag{
			flags.MakeFlag[string]("name", "", nil),
		},
		Name:  "Hello",
		Usage: "",
		Action: func(c *cli.Context) error {
			Hello(
				flags.GetFlag[string](c, "name"),
			)
			return nil
		},
		CtxFlagBuilder: func(c *cli.Context) map[string]any {
			cflags := make(map[string]any)
			cflags["name"] = flags.GetFlag[string](c, "name")
			return cflags
		},
	})
}
