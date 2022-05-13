package eval

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/game"
)

func Eval(st *State) int {
	if st.WK == 0 {
		return -5000
	}
	if st.BK == 0 {
		return 5000
	}

	
}

func CountPiece(piece uint64) float64 {
	var res float64
	for piece != 0 {
		res += 1
		piece &= (piece-1)
	}
	return res
}

func SquaresSeenDiff(st *State) float64 {
	for _, k := 
}

