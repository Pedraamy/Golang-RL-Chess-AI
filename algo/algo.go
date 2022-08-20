package algo

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/eval"
	"math/rand"
    "time"
	//"fmt"
)


func BestMove (st *state.State, depth int) *state.Move {
	if st.White == 1 {
		return BestMoveWhite(st, depth)
	} else {
		return BestMoveBlack(st, depth)
	}
}

func BestMoveWhite (st *state.State, depth int) *state.Move {
	var alpha int = -100001
	var best *state.Move
	var res int = -100001
	var curr int
	var ns *state.State
	captures, moves := st.GetAllMoves()
	total := append(captures, moves...)
	for _, m := range total {
		ns = st.StateFromMove(m)
		curr = MiniMaxBlack(ns, alpha, 100001, depth-1)
		if curr > res {
			res = curr
			best = m
			alpha = res
		}
	}
	return best

}
func BestMoveBlack (st *state.State, depth int) *state.Move {
	var beta int = 100001
	var best *state.Move
	var res int = 100001
	var curr int
	var ns *state.State
	captures, moves := st.GetAllMoves()
	total := append(captures, moves...)
	for _, m := range total {
		ns = st.StateFromMove(m)
		curr = MiniMaxWhite(ns, -100001, beta, depth-1)
		if curr < res {
			res = curr
			best = m
			beta = res
		}
	}
	return best
}

func MiniMaxWhite(st *state.State, alpha int, beta int, depth int) int {
	if st.WK == 0 {
		return -100000
	}
	if depth == 0{
		return eval.Eval(st)
	}
	var res int = -100001
	var curr int
	var ns *state.State
	captures, moves := st.GetAllMoves()
	total := append(captures, moves...)
	if len(total) == 0 {
		return 0
	}
	for _, m := range total {
		ns = st.StateFromMove(m)
		curr = MiniMaxBlack(ns, alpha, beta, depth-1)
		res = Max2(res, curr)
		alpha = Max2(alpha, res)
		if alpha >= beta {
			return res
		}
	}
	return res
	
}

func MiniMaxBlack(st *state.State, alpha int, beta int, depth int) int {
	if st.BK == 0 {
		return 100000
	}
	if depth == 0{
		return eval.Eval(st)
	}
	var res int = 100001
	var curr int
	var ns *state.State
	captures, moves := st.GetAllMoves()
	total := append(captures, moves...)
	if len(total) == 0 {
		return 0
	}
	for _, m := range total {
		ns = st.StateFromMove(m)
		curr = MiniMaxWhite(ns, alpha, beta, depth-1)
		res = Min2(res, curr)
		beta = Min2(beta, res)
		if beta <= alpha {
			return res
		}
	}
	return res	
}

func QuiescenceSearch(st *state.State, alpha int, beta int) int {
	if st.White == 1 {
		return QuiescenceSearchWhite(st, alpha, beta, 0)
	} else {
		return QuiescenceSearchBlack(st, alpha, beta, 0)
	}
}

func QuiescenceSearchWhite(st *state.State, alpha int, beta int, depth int) int {
	if st.WK == 0 {
		return -100000
	}
	captures, _ := st.GetAllMoves()
	if len(captures) == 0 || depth == 2{
		return eval.Eval(st)
	}
	var res int = -100001
	var curr int
	var ns *state.State
	for _, c := range captures {
		ns = st.StateFromMove(c)
		curr = QuiescenceSearchBlack(ns, alpha, beta, depth+1)
		res = Max2(res, curr)
		alpha = Max2(alpha, res)
		if alpha >= beta {
			return res
		}
	}
	return res
}

func QuiescenceSearchBlack(st *state.State, alpha int, beta int, depth int) int {
	if st.BK == 0 {
		return 100000
	}
	captures, _ := st.GetAllMoves()
	if len(captures) == 0 || depth == 6{
		return eval.Eval(st)
	}
	var res int = 100001
	var curr int
	var ns *state.State
	for _, c := range captures {
		ns = st.StateFromMove(c)
		curr = QuiescenceSearchWhite(ns, alpha, beta, depth+1)
		res = Min2(res, curr)
		beta = Min2(beta, res)
		if beta <= alpha {
			return res
		}
	}
	return res
}




/* func MiniMaxWhite(st *state.State, alpha float64, beta float64, depth int) float64 {
	if depth == 0{
		return Eval(st)
	}
	var res float64 = -5001
	var curr *state.State
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white)
		for _, m := range moves {
			mv := move.NewMove(0, st.WK, k, m)
			curr := st.StateFromMoveWhite(mv)
			res = Max2(res, MiniMaxBlack(curr, alpha, beta, depth-1))
			alpha = Max2(alpha, res)
			if alpha >= beta {
				return res
			}
		}
	}
	
	//Queens
	queens := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveWhite(1, queens, q, m)
			res = Max2(res, MiniMaxBlack(curr, alpha, beta, depth-1))
			alpha = Max2(alpha, res)
			if alpha >= beta {
				return res
			}
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.WR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveWhite(2, rooks, r, m)
			res = Max2(res, MiniMaxBlack(curr, alpha, beta, depth-1))
			alpha = Max2(alpha, res)
			if alpha >= beta {
				return res
			}
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.WB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveWhite(3, bishops, b, m)
			res = Max2(res, MiniMaxBlack(curr, alpha, beta, depth-1))
			alpha = Max2(alpha, res)
			if alpha >= beta {
				return res
			}
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.WN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveWhite(4, knights, n, m)
			res = Max2(res, MiniMaxBlack(curr, alpha, beta, depth-1))
			alpha = Max2(alpha, res)
			if alpha >= beta {
				return res
			}
		}
	}
	//Pawn
	rooks := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveWhite(5, pawns, p, m)
			res = Max2(res, MiniMaxBlack(curr, alpha, beta, depth-1))
			alpha = Max2(alpha, res)
			if alpha >= beta {
				return res
			}
		}
	}
	return res
}



func MiniMaxBlack(st *State, depth int) {
	if depth == 0{
		return eval(st)
	}
	var res float64 = 5001
	var curr *State
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveBlack(0, kings, k, m)
			res = Min2(res, MiniMaxWhite(curr, alpha, beta, depth-1))
			beta = Min2(beta, res)
			if beta <= alpha {
				return res
			}
		}
	}
	
	//Queens
	queens := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveBlack(1, queens, q, m)
			res = Min2(res, MiniMaxWhite(curr, alpha, beta, depth-1))
			beta = Min2(beta, res)
			if beta <= alpha {
				return res
			}
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.WR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveBlack(2, rooks, r, m)
			res = Min2(res, MiniMaxWhite(curr, alpha, beta, depth-1))
			beta = Min2(beta, res)
			if beta <= alpha {
				return res
			}
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.WB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveBlack(3, bishops, b, m)
			res = Min2(res, MiniMaxWhite(curr, alpha, beta, depth-1))
			beta = Min2(beta, res)
			if beta <= alpha {
				return res
			}
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.WN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveBlack(4, knights, n, m)
			res = Min2(res, MiniMaxWhite(curr, alpha, beta, depth-1))
			beta = Min2(beta, res)
			if beta <= alpha {
				return res
			}
		}
	}
	//Pawn
	rooks := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveBlack(5, pawns, p, m)
			res = Min2(res, MiniMaxWhite(curr, alpha, beta, depth-1))
			beta = Min2(beta, res)
			if beta <= alpha {
				return res
			}
		}
	}
	return res
}*/

/* 
func Eval(st *state.State) float64 {
	if st.WK == 0 {
		return -5000
	}
	if st.BK == 0 {
		return 5000
	}

	var quotient float64 = 7.5
	var res float64
	black := st.AllBlackPieces()
	white := st.AllWhitePieces()
	wqs := pieces.GetPositionsFromBoard(st.WQ)
	
	res += float64(len(wqs)*9)
	for _, q := range wqs {
		res += float64(len(pieces.QueenMoves(q, white, black)))/quotient
	}
	bqs := pieces.GetPositionsFromBoard(st.BQ)
	res -= float64(len(bqs)*9)
	for _, q := range bqs {
		res -= float64(len(pieces.QueenMoves(q, black, white)))/quotient
	}
	wrs := pieces.GetPositionsFromBoard(st.WR)
	res += float64(len(wrs)*5)
	for _, r := range wrs {
		res += float64(len(pieces.RookMoves(r, white, black)))/quotient
	}
	brs := pieces.GetPositionsFromBoard(st.BR)
	res -= float64(len(brs)*5)
	for _, r := range brs {
		res -= float64(len(pieces.RookMoves(r, black, white)))/quotient
	}
	wbs := pieces.GetPositionsFromBoard(st.WB)
	res += float64(len(wbs))*float64(3.1)
	for _, b := range wbs {
		res += float64(len(pieces.BishopMoves(b, white, black)))/quotient
	}
	bbs := pieces.GetPositionsFromBoard(st.BB)
	res -= float64(len(bbs))*float64(3.1)
	for _, b := range bbs {
		res -= float64(len(pieces.BishopMoves(b, black, white)))/quotient
	}
	wns := pieces.GetPositionsFromBoard(st.WN)
	res += float64(len(wns))*float64(3)
	for _, n := range wns {
		res += float64(len(pieces.KnightMoves(n, white)))/quotient
	}
	bns := pieces.GetPositionsFromBoard(st.BN)
	res -= float64(len(bns))*float64(3)
	for _, n := range bns {
		res -= float64(len(pieces.KnightMoves(n, black)))/quotient
	}
	wps := pieces.GetPositionsFromBoard(st.WP)
	res += float64(len(wps))
	for _, p := range wps {
		res += float64(len(pieces.PawnMoves(p, white, black, 1)))/quotient
	}
	bps := pieces.GetPositionsFromBoard(st.BP)
	res -= float64(len(bps))
	for _, p := range bps {
		res -= float64(len(pieces.PawnMoves(p, black, white, 0)))/quotient
	}
	return res

} */

func RandomMove(st *state.State) *state.Move {
	captures, moves := st.GetAllMoves()
	total := append(captures, moves...)
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(total))
	return total[idx]

}

/* func GetAllMovesWhite(st *State) []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white, black)
		for _, m := range moves {
			res = append(res, move.NewMove(0, kings, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, white, black)
		for _, m := range moves {
			res = append(res, move.NewMove(1, queens, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.WR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, white, black)
		for _, m := range moves {
			res = append(res, move.NewMove(2, rooks, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.WB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, white, black)
		for _, m := range moves {
			res = append(res, move.NewMove(3, bishops, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.WN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, white, black)
		for _, m := range moves {
			res = append(res, move.NewMove(4, knights, n, m))
		}
	}
	//Pawns
	rooks := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, white, black)
		for _, m := range moves {
			res = append(res, move.NewMove(5, pawns, p, m))
		}
	}

	return res
}

func GetAllMovesBlack(st *State) []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.BK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, black, white)
		for _, m := range moves {
			res = append(res, move.NewMove(0, kings, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.BQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, black, white)
		for _, m := range moves {
			res = append(res, move.NewMove(1, queens, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.BR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, black, white)
		for _, m := range moves {
			res = append(res, move.NewMove(2, rooks, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.BB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, black, white)
		for _, m := range moves {
			res = append(res, move.NewMove(3, bishops, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.BN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, black, white)
		for _, m := range moves {
			res = append(res, move.NewMove(4, knights, n, m))
		}
	}
	//Pawns
	rooks := pieces.GetPositionsFromBoard(st.BP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, black, white)
		for _, m := range moves {
			res = append(res, move.NewMove(5, pawns, p, m))
		}
	}

	return res
} */

func Max2(f1 int, f2 int) int {
	if f2 > f1 {
		return f2
	} else {
		return f1
	}
}

func Min2(f1 int, f2 int) int {
	if f2 < f1 {
		return f2
	} else {
		return f1
	}
}

/* func CountPiece(piece uint64) float64 {
	var res float64
	for piece != 0 {
		res += 1
		piece &= (piece-1)
	}
	return res
}

func SquaresSeenDiff(st *State, white uint64, black uint64) float64 {
	var res float64
	wqs := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range wqs {
		res += len(pieces.QueenMoves(q, same, opp))
	}
	bqs := pieces.GetPositionsFromBoard(st.BQ)
	for _, q := range bqs {
		res += len(pieces.)
	}

	
}
 */
