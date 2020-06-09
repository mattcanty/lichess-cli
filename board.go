package main

import (
	"strings"
)

func getRowStrings(fen string, playingWhite bool) [8]string {
	var rows [8]string

	if playingWhite {
		fen = reverse(fen)
	}

	fenRows := strings.Split(fen, "/")

	for i, row := range fenRows {
		row = strings.ReplaceAll(row, "r", "♖ ")
		row = strings.ReplaceAll(row, "n", "♘ ")
		row = strings.ReplaceAll(row, "b", "♗ ")
		row = strings.ReplaceAll(row, "q", "♕ ")
		row = strings.ReplaceAll(row, "k", "♔ ")
		row = strings.ReplaceAll(row, "p", "♙ ")
		row = strings.ReplaceAll(row, "R", "♜ ")
		row = strings.ReplaceAll(row, "N", "♞ ")
		row = strings.ReplaceAll(row, "B", "♝ ")
		row = strings.ReplaceAll(row, "Q", "♛ ")
		row = strings.ReplaceAll(row, "K", "♚ ")
		row = strings.ReplaceAll(row, "P", "♟ ")

		row = strings.ReplaceAll(row, "1", "- ")
		row = strings.ReplaceAll(row, "2", "- - ")
		row = strings.ReplaceAll(row, "3", "- - - ")
		row = strings.ReplaceAll(row, "4", "- - - - ")
		row = strings.ReplaceAll(row, "5", "- - - - - ")
		row = strings.ReplaceAll(row, "6", "- - - - - - ")
		row = strings.ReplaceAll(row, "7", "- - - - - - - ")
		row = strings.ReplaceAll(row, "8", "- - - - - - - - ")

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
