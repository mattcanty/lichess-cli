package main

import (
	"github.com/alecthomas/kong"
)

type context struct {
	Debug bool
}

var cli struct {
	Debug         bool   `help:"Enable debug mode."`
	LichessAPIKey string `help:"Lichess API key"`

	AsciiMode   bool         `help:"Use Ascii characters instead of unicode chess pieces"`
	ColorBoard  string       `help:"Color of the board. Options: default, black_and_white, blue, cyan, green, magenta, none, red, yellow" default:"default"`
	ColorLegend string       `help:"Color of the legend. Options: none, default" default:"default"`
	ColorPieces string       `help:"Color of the pieces. Options: default, black_and_white, none" default:"default"`
	HideLegend  bool         `help:"Hide legend on the board"`
	Games       gamesCmd     `cmd:"" help:"Lists all active games"`
	G           gamesCmd     `cmd:"" help:"Lists all active games"`
	Play        playCmd      `cmd:"" help:"Play a move in a game."`
	P           playCmd      `cmd:"" help:"Play a move in a game."`
	NAI         newAIGameCmd `cmd:"" help:"Create a game against the AI"`
}

func main() {
	ctx := kong.Parse(&cli, kong.Configuration(kong.JSON, "~/.config/lichess-cli/config.json"))

	err := ctx.Run(&context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
