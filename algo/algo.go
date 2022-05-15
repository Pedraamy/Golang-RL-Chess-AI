package algo

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
)

func BestMove {

}

func MiniMaxWhite(st *State, alpha float64, beta float64, depth int) float64 {
	if depth == 0{
		return eval(st)
	}
	var res float64 = -5001
	var curr *State
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white, black)
		for _, m := range moves {
			curr := st.StateFromMoveWhite(0, kings, k, m)
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
}


func Eval(st *State) float64 {
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
	
	res += len(wqs)*9
	for _, q := range wqs {
		res += len(pieces.QueenMoves(q, white, black))/quotient
	}
	bqs := pieces.GetPositionsFromBoard(st.BQ)
	res -= len(bqs)*9
	for _, q := range bqs {
		res -= len(pieces.QueenMoves(q, black, white))/quotient
	}
	wrs := pieces.GetPositionsFromBoard(st.WR)
	res += len(wrs)*5
	for _, r := range wrs {
		res += len(pieces.RookMoves(r, white, black))/quotient
	}
	brs := pieces.GetPositionsFromBoard(st.BR)
	res -= len(brs)*5
	for _, r := range brs {
		res -= len(pieces.RookMoves(r, black, white))/quotient
	}
	wbs := pieces.GetPositionsFromBoard(st.WB)
	res += len(wbs)*3.1
	for _, b := range wbs {
		res += len(pieces.BishopMoves(b, white, black))/quotient
	}
	bbs := pieces.GetPositionsFromBoard(st.BB)
	res -= len(bbs)*3.1
	for _, b := range bbs {
		res -= len(pieces.BishopMoves(b, black, white))/quotient
	}
	wns := pieces.GetPositionsFromBoard(st.WN)
	res += len(wns)*3
	for _, n := range wns {
		res += len(pieces.KnightMoves(n, white, black))/quotient
	}
	bns := pieces.GetPositionsFromBoard(st.BN)
	res -= len(bns)*3
	for _, n := range bns {
		res -= len(pieces.KnightMoves(n, black, white))/quotient
	}
	wps := pieces.GetPositionsFromBoard(st.WP)
	res += len(wps)
	for _, p := range wps {
		res += len(pieces.PawnMoves(p, white, black))/quotient
	}
	bps := pieces.GetPositionsFromBoard(st.BP)
	res -= len(bps)
	for _, p := range bps {
		res -= len(pieces.PawnMoves(p, black, white))/quotient
	}
	return res

}

func RandomMove(st *State) *State {

	possibleMoves = 
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white, black)
	}
}

func Max2(f1 float64, f2 float64) float64 {
	if f2 > f1 {
		return f2
	} else {
		return f1
	}
}

func Min2(f1 float64, f2 float64) float64 {
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
