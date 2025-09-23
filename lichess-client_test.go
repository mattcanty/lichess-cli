package main

import (
	"testing"
)

func TestEnhanceErrorMessage(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Missing scope",
			expected: "Missing scope: Your API token doesn't have the required permission. Please go to https://lichess.org/account/oauth/token and create a new token with 'Play games with the board API' selected.",
		},
		{
			input:    "Invalid move",
			expected: "Invalid move: Please use algebraic notation like 'e2e4' (from square to square). Example: to move a pawn from e2 to e4, use 'e2e4'.",
		},
		{
			input:    "Not your turn",
			expected: "Not your turn: Please wait for your opponent to make their move before playing again.",
		},
		{
			input:    "Game not found",
			expected: "Game not found: Please check the game ID and ensure you have an active game.",
		},
		{
			input:    "Some other error",
			expected: "Some other error",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		got := enhanceErrorMessage(test.input)
		if got != test.expected {
			t.Errorf("enhanceErrorMessage(%q) = %q, want %q", test.input, got, test.expected)
		}
	}
}