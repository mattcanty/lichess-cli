package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/stretchr/testify/assert"
)

var (
	testConfig = printerConfig{
		colorBoard:  "none",
		colorLegend: "none",
		colorPieces: "none",
		showLegend:  true,
	}
	testConfigColored = printerConfig{
		colorBoard:  "default",
		colorLegend: "default",
		colorPieces: "default",
		showLegend:  true,
	}
	testConfigWithoutLegend = printerConfig{
		colorBoard:  "none",
		colorLegend: "none",
		colorPieces: "none",
		showLegend:  false,
	}
	testGame1 = nowPlaying{
		GameID: "vMB7uwrm",
		Fen:    "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
		Color:  "white",
		Opponent: struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Rating   int    `json:"rating"`
		}{Username: "A.I. level 1"},
	}
	testGame2 = nowPlaying{
		GameID:   "KJWzuxbM",
		Fen:      "r1bqk2r/pppp1ppp/2n1p3/2b5/4P3/2N2N2/PPPP1PPP/R1BQK2R",
		IsMyTurn: true,
		LastMove: "e7c6",
		Color:    "black",
		Opponent: struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Rating   int    `json:"rating"`
		}{Username: "Garry Kasparov"},
	}
)

func Test_getCellColors(t *testing.T) {
	for rowIdx := 0; rowIdx < 8; rowIdx++ {
		for colIdx := 0; colIdx < 8; colIdx++ {
			testName := fmt.Sprintf("[%d][%d]", rowIdx+1, colIdx+1)
			bgColor := text.BgYellow
			if (rowIdx+1+colIdx+1)%2 == 0 { // white cell/background
				bgColor = text.BgHiYellow
			}
			fgColorBlack := text.FgBlack
			fgColorWhite := text.FgHiBlack

			ccPieceBlack := getCellColors(rowIdx, colIdx, PieceRookBlack, testConfigColored)
			assert.Equal(t, text.Colors{bgColor, fgColorBlack}, ccPieceBlack, "black piece on "+testName)

			ccPieceWhite := getCellColors(rowIdx, colIdx, PieceRookWhite, testConfigColored)
			assert.Equal(t, text.Colors{bgColor, fgColorWhite}, ccPieceWhite, "white piece on "+testName)
		}
	}
}

func Test_printGame(t *testing.T) {
	t.Run("game 1", func(t *testing.T) {
		output := printGame(testGame1, testConfig)
		expectedOutput := strings.Join([]string{
			" ♜  ♞  ♝  ♛  ♚  ♝  ♞  ♜  8 ",
			" ♟  ♟  ♟  ♟  ♟  ♟  ♟  ♟  7 ",
			"                         6 ",
			"                         5 ",
			"                         4 ",
			"                         3 ",
			" ♙  ♙  ♙  ♙  ♙  ♙  ♙  ♙  2 ",
			" ♖  ♘  ♗  ♕  ♔  ♗  ♘  ♖  1 ",
			" a  b  c  d  e  f  g  h    ",
		}, "\n")
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("game 1 without legend", func(t *testing.T) {
		output := printGame(testGame1, testConfigWithoutLegend)
		expectedOutput := strings.Join([]string{
			" ♜  ♞  ♝  ♛  ♚  ♝  ♞  ♜ ",
			" ♟  ♟  ♟  ♟  ♟  ♟  ♟  ♟ ",
			"                        ",
			"                        ",
			"                        ",
			"                        ",
			" ♙  ♙  ♙  ♙  ♙  ♙  ♙  ♙ ",
			" ♖  ♘  ♗  ♕  ♔  ♗  ♘  ♖ ",
		}, "\n")
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("game 2", func(t *testing.T) {
		output := printGame(testGame2, testConfig)
		expectedOutput := strings.Join([]string{
			" ♖        ♔  ♕  ♗     ♖  1 ",
			" ♙  ♙  ♙     ♙  ♙  ♙  ♙  2 ",
			"       ♘        ♘        3 ",
			"          ♙              4 ",
			"                ♝        5 ",
			"          ♟     ♞        6 ",
			" ♟  ♟  ♟     ♟  ♟  ♟  ♟  7 ",
			" ♜        ♚  ♛  ♝     ♜  8 ",
			" h  g  f  e  d  c  b  a    ",
		}, "\n")
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("game 2 without legend", func(t *testing.T) {
		output := printGame(testGame2, testConfigWithoutLegend)
		expectedOutput := strings.Join([]string{
			" ♖        ♔  ♕  ♗     ♖ ",
			" ♙  ♙  ♙     ♙  ♙  ♙  ♙ ",
			"       ♘        ♘       ",
			"          ♙             ",
			"                ♝       ",
			"          ♟     ♞       ",
			" ♟  ♟  ♟     ♟  ♟  ♟  ♟ ",
			" ♜        ♚  ♛  ♝     ♜ ",
		}, "\n")
		assert.Equal(t, expectedOutput, output)
	})
}

func Test_printGames(t *testing.T) {
	input := []nowPlaying{
		testGame1,
		testGame2,
	}
	expectedOutput := `┌──────────┬────────────┬────────────────┬───────────┬─────────────────────────────┐
│ ID       │ TURN       │ OPPONENT       │ LAST MOVE │ BOARD                       │
├──────────┼────────────┼────────────────┼───────────┼─────────────────────────────┤
│ vMB7uwrm │ Their Turn │ A.I. level 1   │           │  ♜  ♞  ♝  ♛  ♚  ♝  ♞  ♜  8  │
│          │            │                │           │  ♟  ♟  ♟  ♟  ♟  ♟  ♟  ♟  7  │
│          │            │                │           │                          6  │
│          │            │                │           │                          5  │
│          │            │                │           │                          4  │
│          │            │                │           │                          3  │
│          │            │                │           │  ♙  ♙  ♙  ♙  ♙  ♙  ♙  ♙  2  │
│          │            │                │           │  ♖  ♘  ♗  ♕  ♔  ♗  ♘  ♖  1  │
│          │            │                │           │  a  b  c  d  e  f  g  h     │
├──────────┼────────────┼────────────────┼───────────┼─────────────────────────────┤
│ KJWzuxbM │ Your Turn  │ Garry Kasparov │ e7c6      │  ♖        ♔  ♕  ♗     ♖  1  │
│          │            │                │           │  ♙  ♙  ♙     ♙  ♙  ♙  ♙  2  │
│          │            │                │           │        ♘        ♘        3  │
│          │            │                │           │           ♙              4  │
│          │            │                │           │                 ♝        5  │
│          │            │                │           │           ♟     ♞        6  │
│          │            │                │           │  ♟  ♟  ♟     ♟  ♟  ♟  ♟  7  │
│          │            │                │           │  ♜        ♚  ♛  ♝     ♜  8  │
│          │            │                │           │  h  g  f  e  d  c  b  a     │
└──────────┴────────────┴────────────────┴───────────┴─────────────────────────────┘`

	output := printGames(input, testConfig)
	assert.Equal(t, expectedOutput, output)
	if expectedOutput != output {
		fmt.Println(output)
	}
}

func Test_translateGame(t *testing.T) {
	t.Run("game 1", func(t *testing.T) {
		expectedOutput := [][]Piece{
			{PieceRookBlack, PieceKnightBlack, PieceBishopBlack, PieceQueenBlack, PieceKingBlack, PieceBishopBlack, PieceKnightBlack, PieceRookBlack},
			{PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack},
			{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
			{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
			{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
			{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
			{PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite},
			{PieceRookWhite, PieceKnightWhite, PieceBishopWhite, PieceQueenWhite, PieceKingWhite, PieceBishopWhite, PieceKnightWhite, PieceRookWhite},
		}
		output := translateGame(testGame1)
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("game 2", func(t *testing.T) {
		expectedOutput := [][]Piece{
			{PieceRookWhite, PieceNone, PieceNone, PieceKingWhite, PieceQueenWhite, PieceBishopWhite, PieceNone, PieceRookWhite},
			{PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PieceNone, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite, PiecePawnWhite},
			{PieceNone, PieceNone, PieceKnightWhite, PieceNone, PieceNone, PieceKnightWhite, PieceNone, PieceNone},
			{PieceNone, PieceNone, PieceNone, PiecePawnWhite, PieceNone, PieceNone, PieceNone, PieceNone},
			{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceBishopBlack, PieceNone, PieceNone},
			{PieceNone, PieceNone, PieceNone, PiecePawnBlack, PieceNone, PieceKnightBlack, PieceNone, PieceNone},
			{PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PieceNone, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack, PiecePawnBlack},
			{PieceRookBlack, PieceNone, PieceNone, PieceKingBlack, PieceQueenBlack, PieceBishopBlack, PieceNone, PieceRookBlack},
		}
		output := translateGame(testGame2)
		assert.Equal(t, expectedOutput, output)
	})
}
