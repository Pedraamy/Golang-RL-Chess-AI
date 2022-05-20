package eval

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
	"github.com/Pedraamy/Golang-RL-Chess-AI/values"
	"math/rand"
    "time"
)

func Eval(st *state.State) int {
	var res uint64 = 7
	x := pieces.PosTable[res]
}

func PieceValues(st *state.State) int {
	res := 0
	wpawns := pieces.GetPositionsFromBoard(st.WP)
	res += values.Pawn * len(wpawns)
}