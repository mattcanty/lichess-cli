package main

import (
	"testing"
)

func TestGetRowStrings(t *testing.T) {
	got := getRowStrings("R1BKQBNR/PPP1PPPP/8/3P1N2/8/4bp2/ppp1p1pp/rn1kqbnr", true)

	expected := [8]string{
		"♖ ♘ ♗ ♕ ♔ - ♘ ♖",
		"♙ ♙ - ♙ - ♙ ♙ ♙",
		"- - ♙ ♗ - - - -",
		"- - - - - - - -",
		"- - ♞ - ♟ - - -",
		"- - - - - - - -",
		"♟ ♟ ♟ ♟ - ♟ ♟ ♟",
		"♜ ♞ ♝ ♛ ♚ ♝ - ♜",
	}

	if expected != got {
		t.Errorf("Board was not drawn as expected. Got:\n%s\nExpected:\n%s", got, expected)
	}
}

func TestGetRowStringsPlayingBlack(t *testing.T) {
	got := getRowStrings("R1BKQBNR/PPP1PPPP/8/3P1N2/8/4bp2/ppp1p1pp/rn1kqbnr", false)

	expected := [8]string{
		"♜ - ♝ ♚ ♛ ♝ ♞ ♜",
		"♟ ♟ ♟ - ♟ ♟ ♟ ♟",
		"- - - - - - - -",
		"- - - ♟ - ♞ - -",
		"- - - - - - - -",
		"- - - - ♗ ♙ - -",
		"♙ ♙ ♙ - ♙ - ♙ ♙",
		"♖ ♘ - ♔ ♕ ♗ ♘ ♖",
	}

	if expected != got {
		t.Errorf("Board was not drawn as expected. Got:\n%s\nExpected:\n%s", got, expected)
	}
}
