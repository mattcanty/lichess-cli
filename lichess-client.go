package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type nowPlaying struct {
	FullID   string `json:"fullId"`
	GameID   string `json:"gameId"`
	Fen      string `json:"fen"`
	Color    string `json:"color"`
	LastMove string `json:"lastMove"`
	Variant  struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"variant"`
	Speed    string `json:"speed"`
	Perf     string `json:"perf"`
	Rated    bool   `json:"rated"`
	HasMoved bool   `json:"hasMoved"`
	Opponent struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Rating   int    `json:"rating"`
	} `json:"opponent"`
	IsMyTurn    bool `json:"isMyTurn"`
	SecondsLeft int  `json:"secondsLeft"`
}

type nowPlayingResponse struct {
	NowPlaying []nowPlaying `json:"nowPlaying"`
}

type boardMoveResponse struct {
	Ok    bool
	Error string
}

func get(lichessAPIKey string) ([]nowPlaying, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://lichess.org/api/account/playing", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lichessAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var data nowPlayingResponse

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}
	return data.NowPlaying, nil
}

func post(lichessAPIKey string, gameID string, move string) (string, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://lichess.org//api/board/game/%s/move/%s", gameID, move)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lichessAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var boardMoveResponse boardMoveResponse
	json.Unmarshal(bodyBytes, &boardMoveResponse)

	return boardMoveResponse.Error, nil
}
