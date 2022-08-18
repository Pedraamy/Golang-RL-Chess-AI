package eval

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
	"github.com/Pedraamy/Golang-RL-Chess-AI/values"
)

func Later(st *state.State) int {
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
	//return PieceValues(st)
	return 5
}

func Eval(st *state.State) int {
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

	res += EarlyQueenDev(st)
	res += BishopPairAdv(wbishops, bbishops)

	return res

}

func EarlyQueenDev(st *state.State) int {
	res := 0
	if st.WQ&(1<<4) == 0 {
		if st.WB&(1<<2) != 0 {
			res -= 12
		}
		if st.WB&(1<<5) != 0 {
			res -= 12
		}
		if st.WN&(1<<1) != 0 {
			res -= 12
		}
		if st.WN&(1<<6) != 0 {
			res -= 12
		}
	}
	if st.BQ&(1<<4) == 0 {
		if st.BB&(1<<58) != 0 {
			res += 12
		}
		if st.BB&(1<<61) != 0 {
			res += 12
		}
		if st.BN&(1<<57) != 0 {
			res += 12
		}
		if st.BN&(1<<62) != 0 {
			res += 12
		}
	}
	return res
}

func BishopPairAdv(wbishops []uint64, bbishops []uint64) int {
	res := 0
	if len(wbishops) == 2 {
		res += 60
	}
	if len(bbishops) == 2 {
		res -= 60
	}
	return res
}