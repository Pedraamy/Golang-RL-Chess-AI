package utils

import (
	"strings"
	"strconv"
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/algo"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
)

func BestMoveFromFen (fen string) string {
	boardState := StateFromFen(fen)
	bestMove := algo.BestMove(boardState, 5)
	posTable := pieces.PosTable
	begin := posTable[bestMove.Start]
	end := posTable[bestMove.Start]
	beginStr := strconv.FormatUint(begin, 10)
	endStr := strconv.FormatUint(end, 10)
	stringResponse := beginStr + "," + endStr
	if bestMove.Castle == 1 {
		stringResponse = "K"
	} else if bestMove.Castle == 2{
		stringResponse = "Q"
	}
	return stringResponse
} 


func StateFromFen (fen string) *state.State {
	arr := strings.Split(fen, "/")
	boardString, castles, color := arr[0], arr[1], arr[2]
	board := BoardFromString(boardString)
	return BoardToState(board, castles, color)
}

func BoardToState(board [64]int, castles string, color string) *state.State {
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

func BoardFromString (boardString string) [64]int {
	board := [64]int{}
	for i := 0; i<64; i++ {
		if boardString[i] != 32 {
			board[i] = int(boardString[i])
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