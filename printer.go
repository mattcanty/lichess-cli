package main

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// colorCode controls the colors used for the pieces and the board
type colorCode struct {
	Black text.Colors
	White text.Colors
}

// printerConfig controls how the game boards are printed
type printerConfig struct {
	colorBoard  string
	colorLegend string
	colorPieces string
	ascii       bool
	showLegend  bool
}

func (pc *printerConfig) clean() {
	pc.colorBoard = strings.ToLower(pc.colorBoard)
	pc.colorLegend = strings.ToLower(pc.colorLegend)
	pc.colorPieces = strings.ToLower(pc.colorPieces)
}

// defaults and translation maps
var (
	colorBoardMap = map[string]*colorCode{
		"default":         {Black: text.Colors{text.BgYellow}, White: text.Colors{text.BgHiYellow}},
		"black_and_white": {Black: text.Colors{text.BgHiWhite}, White: text.Colors{text.BgWhite}},
		"blue":            {Black: text.Colors{text.BgBlue}, White: text.Colors{text.BgHiBlue}},
		"cyan":            {Black: text.Colors{text.BgCyan}, White: text.Colors{text.BgHiCyan}},
		"green":           {Black: text.Colors{text.BgGreen}, White: text.Colors{text.BgHiGreen}},
		"magenta":         {Black: text.Colors{text.BgMagenta}, White: text.Colors{text.BgHiMagenta}},
		"none":            {Black: text.Colors{}, White: text.Colors{}},
		"red":             {Black: text.Colors{text.BgRed}, White: text.Colors{text.BgHiRed}},
		"yellow":          {Black: text.Colors{text.BgYellow}, White: text.Colors{text.BgHiYellow}},
	}
	colorLegendMap = map[string]text.Colors{
		"none": {},
		"default": {
			text.Italic,    // Italics may not work in all consoles
			text.FgHiBlack, // HiBlack == Gray
		},
	}
	colorPiecesMap = map[string]*colorCode{
		"default":         {Black: text.Colors{text.FgBlack}, White: text.Colors{text.FgHiBlack}},
		"black_and_white": {Black: text.Colors{text.FgBlack}, White: text.Colors{text.FgWhite}},
		"none":            {Black: text.Colors{}, White: text.Colors{}},
	}
	fenPieceMap = map[rune][]Piece{
		'1': {PieceNone},
		'2': {PieceNone, PieceNone},
		'3': {PieceNone, PieceNone, PieceNone},
		'4': {PieceNone, PieceNone, PieceNone, PieceNone},
		'5': {PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		'6': {PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		'7': {PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		'8': {PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		'B': {PieceBishopWhite},
		'K': {PieceKingWhite},
		'N': {PieceKnightWhite},
		'P': {PiecePawnWhite},
		'Q': {PieceQueenWhite},
		'R': {PieceRookWhite},
		'b': {PieceBishopBlack},
		'k': {PieceKingBlack},
		'n': {PieceKnightBlack},
		'p': {PiecePawnBlack},
		'q': {PieceQueenBlack},
		'r': {PieceRookBlack},
	}
	legendRowWhite      = table.Row{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h "}
	legendRowBlack      = table.Row{" h ", " g ", " f ", " e ", " d ", " c ", " b ", " a "}
	asciiPieceStringMap = map[Piece]string{
		PieceBishopBlack: " B ",
		PieceBishopWhite: " B ",
		PieceKingBlack:   " K ",
		PieceKingWhite:   " K ",
		PieceKnightBlack: " N ",
		PieceKnightWhite: " N ",
		PieceNone:        "   ",
		PiecePawnBlack:   " P ",
		PiecePawnWhite:   " P ",
		PieceQueenBlack:  " Q ",
		PieceQueenWhite:  " Q ",
		PieceRookBlack:   " R ",
		PieceRookWhite:   " R ",
	}
	unicodePieceStringMap = map[Piece]string{
		PieceBishopBlack: " ♝ ",
		PieceBishopWhite: " ♗ ",
		PieceKingBlack:   " ♚ ",
		PieceKingWhite:   " ♔ ",
		PieceKnightBlack: " ♞ ",
		PieceKnightWhite: " ♘ ",
		PieceNone:        "   ",
		PiecePawnBlack:   " ♟ ",
		PiecePawnWhite:   " ♙ ",
		PieceQueenBlack:  " ♛ ",
		PieceQueenWhite:  " ♕ ",
		PieceRookBlack:   " ♜ ",
		PieceRookWhite:   " ♖ ",
	}
)

func getCellColors(rowIdx int, colIdx int, piece Piece, cfg printerConfig) text.Colors {
	cBoard, cPieces := colorBoardMap["default"], colorPiecesMap["default"]
	if colors := colorBoardMap[cfg.colorBoard]; colors != nil {
		cBoard = colors
	}
	if colors := colorPiecesMap[cfg.colorPieces]; colors != nil {
		cPieces = colors
	}

	cellColor := cBoard.Black
	if (rowIdx+1+colIdx+1)%2 == 0 {
		cellColor = cBoard.White
	}
	if piece.isBlack() {
		cellColor = append(cellColor, cPieces.Black...)
	} else {
		cellColor = append(cellColor, cPieces.White...)
	}
	return cellColor
}

func getLegendColors(cfg printerConfig) text.Colors {
	if colors, ok := colorLegendMap[cfg.colorLegend]; ok {
		return colors
	}
	return colorLegendMap["default"]
}

func printGame(game nowPlaying, cfg printerConfig) string {
	t := table.NewWriter()

	// loop through each line in the game map and render each row
	colorLegend := getLegendColors(cfg)
	for rowIdx, row := range translateGame(game) {
		rowColorized := table.Row{}
		for colIdx, col := range row {
			cellColors := getCellColors(rowIdx, colIdx, row[colIdx], cfg)
			if cfg.ascii {
				rowColorized = append(rowColorized, cellColors.Sprint(asciiPieceStringMap[col]))
			} else {
				rowColorized = append(rowColorized, cellColors.Sprint(unicodePieceStringMap[col]))
			}
		}
		if cfg.showLegend {
			if game.Color == "white" {
				rowColorized = append(rowColorized, colorLegend.Sprintf(" %d ", 8-rowIdx))
			} else {
				rowColorized = append(rowColorized, colorLegend.Sprintf(" %d ", rowIdx+1))
			}
		}
		t.AppendRow(rowColorized)
	}
	if cfg.showLegend {
		row := table.Row{}
		var legendRow table.Row
		if game.Color == "white" {
			legendRow = legendRowWhite
		} else {
			legendRow = legendRowBlack
		}
		for _, col := range legendRow {
			row = append(row, colorLegend.Sprint(col))
		}
		t.AppendRow(row)
	}

	// set up the options to not draw any separators to make it look like a real
	// chess board; remove padding as we need continuous coloring
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.Style().Options.SeparateRows = false
	t.Style().Box.PaddingLeft = ""
	t.Style().Box.PaddingRight = ""

	return t.Render()
}

func printGames(nowPlaying []nowPlaying, cfg printerConfig) string {
	cfg.clean()

	t := table.NewWriter()
	t.AppendHeader(table.Row{"ID", "Turn", "Opponent", "Last Move", "Board"})
	for _, game := range nowPlaying {
		turn := "Their Turn"
		if game.IsMyTurn {
			turn = "Your Turn"
		}
		board := printGame(game, cfg)

		t.AppendRow(table.Row{game.GameID, turn, game.Opponent.Username, game.LastMove, board})
	}
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true

	return t.Render()
}

func printMoveMessage(move string, message string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Attempted Move", "Message"})
	t.AppendRow([]interface{}{move, message})
	t.SetStyle(table.StyleBold)
	t.Render()
}

func printNewGameId(gameId string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Game Id"})
	t.AppendRow([]interface{}{gameId})
	t.SetStyle(table.StyleBold)
	t.Render()
}

func translateGame(game nowPlaying) [][]Piece {
	fen := strings.Split(game.Fen, " ")[0]
	var rsp [][]Piece

	if game.Color == "white" {
		for _, row := range strings.Split(fen, "/") {
			var rspRow []Piece
			for _, col := range row {
				rspRow = append(rspRow, fenPieceMap[col]...)
			}
			rsp = append(rsp, rspRow)
		}
	} else {
		fens := strings.Split(fen, "/")
		for i, _ := range fens {
			row := fens[len(fens)-1-i]
			var rspRow []Piece
			for j, _ := range row {
				col := []rune(row)[len(row)-1-j]
				rspRow = append(rspRow, fenPieceMap[col]...)
			}
			rsp = append(rsp, rspRow)
		}
	}
	return rsp
}
