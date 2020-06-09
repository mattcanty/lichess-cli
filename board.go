package main

import (
	"strings"
)

var fenPieceMappings = map[string]string{
	"r": "♖ ",
	"n": "♘ ",
	"b": "♗ ",
	"q": "♕ ",
	"k": "♔ ",
	"p": "♙ ",
	"R": "♜ ",
	"N": "♞ ",
	"B": "♝ ",
	"Q": "♛ ",
	"K": "♚ ",
	"P": "♟ ",
	"1": strings.Repeat("- ", 1),
	"2": strings.Repeat("- ", 2),
	"3": strings.Repeat("- ", 3),
	"4": strings.Repeat("- ", 4),
	"5": strings.Repeat("- ", 5),
	"6": strings.Repeat("- ", 6),
	"7": strings.Repeat("- ", 7),
	"8": strings.Repeat("- ", 8),
}

func getRowStrings(fen string, playingWhite bool) [8]string {
	var rows [8]string

	if playingWhite {
		fen = reverse(fen)
	}

	fenRows := strings.Split(fen, "/")

	for i, row := range fenRows {
		for k, v := range fenPieceMappings {
			row = strings.ReplaceAll(row, k, v)
		}

		rows[i] = strings.Trim(row, " ")
	}

	return rows
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
