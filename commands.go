package main

import (
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

type gamesCmd struct {
}

func (r *gamesCmd) Run(ctx *context) error {
	nowPlaying, err := get(cli.LichessAPIKey)
	if err != nil {
		return err
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "My Turn", "Opponent", "Last Move", "Colour"})
	for i, game := range nowPlaying {
		t.AppendRow([]interface{}{i, game.IsMyTurn, game.Opponent.Username, game.LastMove, game.Color})
	}
	t.Render()

	return nil
}

type viewCmd struct {
	GameNumber string `arg name:"game-number" help:"The game number." type:"int"`
}

func (r *viewCmd) Run(ctx *context) error {
	gameNumber, err := strconv.Atoi(r.GameNumber)
	if err != nil {
		return err
	}

	nowPlaying, err := get(cli.LichessAPIKey)
	if err != nil {
		return err
	}
	game := nowPlaying[gameNumber]
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "My Turn", "Opponent", "Last Move", "Colour"})
	t.AppendRow([]interface{}{gameNumber, game.IsMyTurn, game.Opponent.Username, game.LastMove, game.Color})
	t.Render()

	drawBoard(game.Fen, game.Color == "black")
	return nil
}

type playCmd struct {
	GameNumber string `arg name:"game-number" help:"The game number." type:"int"`
	Move       string `arg name:"move" help:"The move." type:"string"`
}

func (r *playCmd) Run(ctx *context) error {
	gameNumber, err := strconv.Atoi(r.GameNumber)
	if err != nil {
		return err
	}

	nowPlaying, err := get(cli.LichessAPIKey)
	if err != nil {
		return err
	}
	game := nowPlaying[gameNumber]

	err = post(cli.LichessAPIKey, game.FullID, r.Move)
	if err != nil {
		return err
	}

	return nil
}
