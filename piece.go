package main

import "strings"

// Piece defines a single Chess Piece
type Piece string

const (
	PieceBishopBlack Piece = "BishopBlack"
	PieceBishopWhite Piece = "BishopWhite"
	PieceKingBlack   Piece = "KingBlack"
	PieceKingWhite   Piece = "KingWhite"
	PieceKnightBlack Piece = "KnightBlack"
	PieceKnightWhite Piece = "KnightWhite"
	PieceNone        Piece = "None"
	PiecePawnBlack   Piece = "PawnBlack"
	PiecePawnWhite   Piece = "PawnWhite"
	PieceQueenBlack  Piece = "QueenBlack"
	PieceQueenWhite  Piece = "QueenWhite"
	PieceRookBlack   Piece = "RookBlack"
	PieceRookWhite   Piece = "RookWhite"
)

func (p Piece) isBlack() bool {
	return strings.HasSuffix(string(p), "Black")
}

func (p Piece) isWhite() bool {
	return strings.HasSuffix(string(p), "White")
}
