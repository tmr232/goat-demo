package main

import (
	"github.com/tmr232/goat"
	"github.com/tmr232/goat/flags"
	"github.com/urfave/cli/v2"
)

func init() {
	goat.Register(Greet, goat.RunConfig{
		Flags: []cli.Flag{
			flags.MakeFlag[string]("name", "The name to greet", nil),
			flags.MakeFlag[bool]("goodbye", "Say goodbye instead of hello", nil),
			flags.MakeFlag[int]("times", "Repeat greeting multiple times", 1),
			flags.MakeFlag[*string]("extra", "Optionally add an extra `message`", nil),
		},
		Name:  "Greet",
		Usage: "greets the given name.",
		Action: func(c *cli.Context) error {
			Greet(
				flags.GetFlag[string](c, "name"),
				flags.GetFlag[bool](c, "goodbye"),
				flags.GetFlag[int](c, "times"),
				flags.GetFlag[*string](c, "extra"),
			)
			return nil
		},
		CtxFlagBuilder: func(c *cli.Context) map[string]any {
			cflags := make(map[string]any)
			cflags["name"] = flags.GetFlag[string](c, "name")
			cflags["goodbye"] = flags.GetFlag[bool](c, "goodbye")
			cflags["times"] = flags.GetFlag[int](c, "times")
			cflags["extra"] = flags.GetFlag[*string](c, "extra")
			return cflags
		},
	})
}
