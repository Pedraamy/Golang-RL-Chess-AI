package utils

import (
	"strings"
)


func TokenizeFen(fen string) []string {
	tokens := strings.Split(fen, " ")
	return tokens
}

func FenToArr (fen string) {
	//tokens := strings.Split(fen, " ")

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
	return board
}


func RotateArr180(board [64]int) {
	left := 0
	right := 63
	for left < right {
		board[left], board[right] = board[right], board[left]
		left++
		right--
	}
}