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
		Name:  "hi",
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

	goat.Register(Goodbye, goat.RunConfig{
		Flags: []cli.Flag{
			flags.MakeFlag[string]("name", "", nil),
			flags.MakeFlag[bool]("formal", "", nil),
		},
		Name:  "bye",
		Usage: "",
		Action: func(c *cli.Context) error {
			Goodbye(
				flags.GetFlag[string](c, "name"),
				flags.GetFlag[bool](c, "formal"),
			)
			return nil
		},
		CtxFlagBuilder: func(c *cli.Context) map[string]any {
			cflags := make(map[string]any)
			cflags["name"] = flags.GetFlag[string](c, "name")
			cflags["formal"] = flags.GetFlag[bool](c, "formal")
			return cflags
		},
	})
}
