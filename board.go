package main

import (
	"fmt"
	"log"
	"strconv"
)

func drawBoard(fen string, playingWhite bool) {
	if playingWhite {
		fen = reverse(fen)
	}
	for _, c := range fen {
		s := fmt.Sprintf("%c", c)
		switch s {
		case "r":
			fmt.Print("♖ ")
		case "n":
			fmt.Print("♘ ")
		case "b":
			fmt.Print("♗ ")
		case "q":
			fmt.Print("♕ ")
		case "k":
			fmt.Print("♔ ")
		case "p":
			fmt.Print("♙ ")
		case "R":
			fmt.Print("♜ ")
		case "N":
			fmt.Print("♞ ")
		case "B":
			fmt.Print("♝ ")
		case "Q":
			fmt.Print("♛ ")
		case "K":
			fmt.Print("♚ ")
		case "P":
			fmt.Print("♟ ")
		case "/":
			fmt.Println("")
		case "1", "2", "3", "4", "5", "6", "7", "8":
			count, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}
			for i := 0; i < count; i++ {
				fmt.Print("- ")
			}
		default:
			log.Fatalf("Unexpected string %s\n", s)
		}
	}

	fmt.Println("")
	fmt.Println("")
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
