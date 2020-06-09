package main

import (
	"strings"

	"github.com/pkg/errors"
)

type gamesCmd struct {
}

func (r *gamesCmd) Run(ctx *context) error {
	nowPlaying, err := get(cli.LichessAPIKey)
	if err != nil {
		return err
	}

	printGames(nowPlaying)

	return nil
}

type playCmd struct {
	GameIdPrefix string `arg name:"game-id-prefix" help:"The game ID." type:"string"`
	Move         string `arg name:"move" help:"The move." type:"string"`
}

func (r *playCmd) Run(ctx *context) error {
	nowPlaying, err := get(cli.LichessAPIKey)
	if err != nil {
		return err
	}
	gameFullID, err := getGameFullId(nowPlaying, r.GameIdPrefix)

	message, err := post(cli.LichessAPIKey, gameFullID, r.Move)

	if err != nil {
		return err
	}

	if message != "" {
		printMoveMessage(r.Move, message)
	}

	return nil
}

func getGameFullId(nowPlaying []nowPlaying, gameIDPrefix string) (string, error) {
	for _, game := range nowPlaying {
		if strings.HasPrefix(game.FullID, gameIDPrefix) {
			return game.FullID, nil
		}
	}
	return "", errors.Errorf("Unable to find game with ID prefixed with: '%s'", gameIDPrefix)
}
