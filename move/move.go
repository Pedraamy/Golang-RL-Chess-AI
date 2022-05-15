package move

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"

)

type Move struct {
	Name uint8
	Piece uint64
	Start uint64
	End uint64
}

func NewMove(name uint8, piece uint64, start uint64, end uint64) *Move {
	return &Move{name, piece, start, end}
}

func GetAllMovesWhite(st *State) []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white, black)
		for _, m := range moves {
			res = append(res, NewMove(0, st.WK, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, white, black)
		for _, m := range moves {
			res = append(res, NewMove(1, st.WQ, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.WR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, white, black)
		for _, m := range moves {
			res = append(res, NewMove(2, st.WR, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.WB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, white, black)
		for _, m := range moves {
			res = append(res, NewMove(3, st.WB, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.WN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, white, black)
		for _, m := range moves {
			res = append(res, NewMove(4, st.WN, n, m))
		}
	}
	//Pawns
	pawns := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, white, black)
		for _, m := range moves {
			res = append(res, NewMove(5, st.WP, p, m))
		}
	}

	return res
}

func GetAllMovesBlack(st *State) []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white, black)
		for _, m := range moves {
			res = append(res, NewMove(0, st.WK, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.WQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, white, black)
		for _, m := range moves {
			res = append(res, NewMove(1, st.WQ, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.WR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, white, black)
		for _, m := range moves {
			res = append(res, NewMove(2, st.WR, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.WB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, white, black)
		for _, m := range moves {
			res = append(res, NewMove(3, st.WB, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.WN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, white, black)
		for _, m := range moves {
			res = append(res, NewMove(4, st.WN, n, m))
		}
	}
	//Pawns
	pawns := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.RookMoves(p, white, black)
		for _, m := range moves {
			res = append(res, NewMove(5, st.WP, p, m))
		}
	}

	return res
}