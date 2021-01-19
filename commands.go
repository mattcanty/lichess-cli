package main

import (
	"errors"
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
		colorBoard:  cli.ColorBoard,
		colorLegend: cli.ColorLegend,
		colorPieces: cli.ColorPieces,
		ascii:       cli.AsciiMode,
		showLegend:  !cli.HideLegend,
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

type newAIGameCmd struct {
	Level int    `arg:"" name:"level" help:"The AI level, 1-12" type:"int"`
	Color string `arg:"" name:"color" help:"The color you want to be, white or black" type:"string"`
}

func (r *newAIGameCmd) Run(ctx *context) error {
	if r.Level < 1 || r.Level > 12 {
		return errors.New("Level must be >= 1 and <= 12")
	}
	if r.Color != "white" && r.Color != "black" {
		return errors.New("Color must be either \"white\" or \"black\"")
	}

	gameId, err := postChallengeAI(cli.LichessAPIKey, r.Level, r.Color)
	if err != nil {
		return err
	}

	printNewGameId(gameId)

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
