package main

import (
	"bytes"
	"encoding/json"
	"errors"
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

type challengeAIBody struct {
	Level int    `json:"level"`
	Color string `json:"color"`
}

type user struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Patron bool   `json:"patron"`
	Id     string `json:"id"`
}

type ratedUser struct {
	User       user
	Rating     int `json:"rating"`
	RatingDiff int `json:"ratingDiff`
}

type gamePlayers struct {
	White ratedUser `json:"white"`
	Black ratedUser `json:"black"`
}

type opening struct {
	Eco  string `json:"eco"`
	Name string `json:"name"`
	Ply  int    `json:"ply"`
}

type clock struct {
	Initial   int `json:"initial"`
	Increment int `json:"increment"`
	TotalTime int `json:"totalTime"`
}

type game struct {
	Id         string      `json:"id"`
	Rated      bool        `json:"rated"`
	Variant    string      `json:"variant"`
	Speed      string      `json:"speed"`
	Perf       string      `json:"perf"`
	CreatedAt  int         `json:"createdAt"`
	LastMoveAt int         `json:"lastMoveAt"`
	Status     string      `json:"status"`
	Players    gamePlayers `json:"players"`
	Opening    opening     `json:"opening"`
	Moves      string      `json:"moves"`
	Clock      clock       `json:"clock"`
}

type genericErrorResponse struct {
	Error string
}

func getAccountPlaying(lichessAPIKey string) ([]nowPlaying, error) {
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

func postBoardGameMove(lichessAPIKey string, gameID string, move string) (string, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://lichess.org/api/board/game/%s/move/%s", gameID, move)

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

func postChallengeAI(lichessAPIKey string, level int, color string) (string, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://lichess.org/api/challenge/ai")
	b := challengeAIBody{Level: level, Color: color}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(b)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lichessAPIKey))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		var err genericErrorResponse
		json.Unmarshal(bodyBytes, &err)
		return "", errors.New(err.Error)
	} else {
		var challengeAIResponse game
		json.Unmarshal(bodyBytes, &challengeAIResponse)
		return challengeAIResponse.Id, nil
	}
}
