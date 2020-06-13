package main

import (
	"fmt"
	"strings"
)

type gamesCmd struct {
}

func (r *gamesCmd) Run(ctx *context) error {
	nowPlaying, err := getAccountPlaying(cli.LichessAPIKey)
	if err != nil {
		return err
	}

	fmt.Println(printGames(nowPlaying, printerConfig{
		colorBoard:  "default",
		colorLegend: "default",
		colorPieces: "default",
		showLegend:  true,
	}))

	return nil
}

type playCmd struct {
	GameIDPrefix string `arg:"" name:"game-id-prefix" help:"The game ID." type:"string"`
	Move         string `arg:"" name:"move" help:"The move." type:"string"`
}

func (r *playCmd) Run(ctx *context) error {
	nowPlaying, err := getAccountPlaying(cli.LichessAPIKey)
	if err != nil {
		return err
	}
	gameFullID, err := getGameFullId(nowPlaying, r.GameIDPrefix)

	message, err := postBoardGameMove(cli.LichessAPIKey, gameFullID, r.Move)

	if err != nil {
		return err
	}

	if message != "" {
		printMoveMessage(r.Move, message)
	}

	return nil
}

func getGameFullId(nowPlaying []nowPlaying, gameIDPrefix string) (string, error) {
	var matches []string
	for _, game := range nowPlaying {
		if strings.HasPrefix(game.FullID, gameIDPrefix) {
			matches = append(matches, game.FullID)
		}
	}

	if len(matches) == 1 {
		return matches[0], nil
	}

	if len(matches) > 1 {
		return "", fmt.Errorf("Prefix '%s' matches multiple game IDs: %s", gameIDPrefix, matches)
	}

	return "", fmt.Errorf("Unable to find game with ID prefixed with: '%s'", gameIDPrefix)
}
