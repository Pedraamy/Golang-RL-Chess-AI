package eval

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
	"github.com/Pedraamy/Golang-RL-Chess-AI/values"
)

func Eval(st *state.State) int {
	/* wpawns := pieces.GetPiecesFromBoard(st.WP)
	bpawns := pieces.GetPiecesFromBoard(st.BP)
	wknights := pieces.GetPiecesFromBoard(st.WN)
	bknights := pieces.GetPiecesFromBoard(st.BN)
	wbishops := pieces.GetPiecesFromBoard(st.WB)
	bbishops := pieces.GetPiecesFromBoard(st.BB)
	wrooks := pieces.GetPiecesFromBoard(st.WR)
	brooks := pieces.GetPiecesFromBoard(st.BR)
	wqueens := pieces.GetPiecesFromBoard(st.WQ)
	bqueens := pieces.GetPiecesFromBoard(st.BQ)
	wking := pieces.GetPiecesFromBoard(st.WK)
	bking := pieces.GetPiecesFromBoard(st.BK) */
	return PieceValues(st)
}

func PieceValues(st *state.State) int {
	res := 0
	wpawns := pieces.GetPiecesFromBoard(st.WP)
	res += values.Pawn * len(wpawns)
	for _, p := range wpawns {
		idx := pieces.PosTable[p]
		res += values.WPawn[idx]
	}
	bpawns := pieces.GetPiecesFromBoard(st.BP)
	res -= values.Pawn * len(bpawns)
	for _, p := range bpawns {
		idx := pieces.PosTable[p]
		res -= values.BPawn[idx]
	}
	wknights := pieces.GetPiecesFromBoard(st.WN)
	res += values.Knight * len(wknights)
	for _, p := range wknights {
		idx := pieces.PosTable[p]
		res += values.WKnight[idx]
	}
	bknights := pieces.GetPiecesFromBoard(st.BN)
	res -= values.Knight * len(bknights)
	for _, p := range bknights {
		idx := pieces.PosTable[p]
		res -= values.BKnight[idx]
	}
	wbishops := pieces.GetPiecesFromBoard(st.WB)
	res += values.Bishop * len(wbishops)
	for _, p := range wbishops {
		idx := pieces.PosTable[p]
		res += values.WBishop[idx]
	}
	bbishops := pieces.GetPiecesFromBoard(st.BB)
	res -= values.Bishop * len(bbishops)
	for _, p := range bbishops {
		idx := pieces.PosTable[p]
		res -= values.BBishop[idx]
	}
	wrooks := pieces.GetPiecesFromBoard(st.WR)
	res += values.Rook * len(wrooks)
	for _, p := range wrooks {
		idx := pieces.PosTable[p]
		res += values.WRook[idx]
	}
	brooks := pieces.GetPiecesFromBoard(st.BR)
	res -= values.Rook * len(brooks)
	for _, p := range brooks {
		idx := pieces.PosTable[p]
		res -= values.BRook[idx]
	}
	wqueens := pieces.GetPiecesFromBoard(st.WQ)
	res += values.Queen * len(wqueens)
	for _, p := range wqueens {
		idx := pieces.PosTable[p]
		res += values.WQueen[idx]
	}
	bqueens := pieces.GetPiecesFromBoard(st.BQ)
	res -= values.Queen * len(bqueens)
	for _, p := range bqueens {
		idx := pieces.PosTable[p]
		res -= values.BQueen[idx]
	}
	wking := pieces.GetPiecesFromBoard(st.WK)
	for _, p := range wking {
		idx := pieces.PosTable[p]
		res += values.WKing[idx]
	}
	bking := pieces.GetPiecesFromBoard(st.BK)
	for _, p := range bking {
		idx := pieces.PosTable[p]
		res -= values.BKing[idx]
	}
	return res

}