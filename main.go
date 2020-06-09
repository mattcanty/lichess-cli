package main

import (
	"github.com/alecthomas/kong"
)

type context struct {
	Debug bool
}

var cli struct {
	Debug         bool   `help:"Enable debug mode."`
	LichessAPIKey string `help:"Enable debug mode."`

	Games gamesCmd `cmd help:"Lists all active games"`
	Play  playCmd  `cmd help:"Play a move in a game."`
}

func main() {
	ctx := kong.Parse(&cli, kong.Configuration(kong.JSON, "~/.config/lichess-cli/config.json"))

	err := ctx.Run(&context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
