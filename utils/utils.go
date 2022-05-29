package utils

import (
	"strings"
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
)


func TokenizeFen(fen string) []string {
	tokens := strings.Split(fen, " ")
	return tokens
}

func FenToArr (fen string) {
	//tokens := strings.Split(fen, " ")

}

func StateFromFen(board [64]int, color string, castles string) *state.State {
	var (
		White uint8 = 0
		WK uint64 = 0
		WQ uint64 = 0
		WR uint64 = 0
		WB uint64 = 0
		WN uint64 = 0
		WP uint64 = 0
		BK uint64 = 0
		BQ uint64 = 0
		BR uint64 = 0
		BB uint64 = 0
		BN uint64 = 0
		BP uint64 = 0
		CastleWK uint8 = 0
		CastleWQ uint8 = 0
		CastleBK uint8 = 0
		CastleBQ uint8 = 0
		JustCastled uint8 = 0
	)
	if color == "w" {
		White = 1
	}
	for i := 0; i<len(castles); i++ {
		switch castles[i] {
		case 'K':
			CastleWK = 1
		case 'Q':
			CastleWQ = 1
		case 'k':
			CastleBK = 1
		case 'q':
			CastleBQ = 1
		}
	}
	for i := 0; i<64; i++ {
		switch board[i] {
		case 80:
			WP |= 1<<i
		case 112:
			BP |= 1<<i
		case 78:
			WN |= 1<<i
		case 110:
			BN |= 1<<i
		case 66:
			WB |= 1<<i
		case 98:
			BB |= 1<<i
		case 82:
			WR |= 1<<i
		case 114:
			BR |= 1<<i
		case 81:
			WQ |= 1<<i
		case 113:
			BQ |= 1<<i
		case 75:
			WK |= 1<<i
		case 107:
			BK |= 1<<i
		}
	}
	return &state.State{White, WK, WQ, WR, WB, WN, WP, BK, BQ, BR, BB, BN, BP, CastleWK, CastleWQ, CastleBK, CastleBQ, JustCastled}

}

func BoardFromString (s string) [64]int {
	board := [64]int{}
	pos := 0
	for i := 0; i < len(s); i++ {
		curr := int(s[i])
		if curr == 47 {
			continue
		} else if curr >= 48 && curr <= 57 {
			pos += curr
		} else {
			board[pos] = curr
			pos++
		}
	}
	board = RotateArr180(board)
	return board
}


func RotateArr180(board [64]int) [64]int {
	left := 0
	right := 63
	for left < right {
		board[left], board[right] = board[right], board[left]
		left++
		right--
	}
	return board
}