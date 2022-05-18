package state

import (
	"math"
	"fmt"
	"strconv"
	"github.com/jinzhu/copier"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
)

var (
	posTable = PosTable()
)

type State struct {
	White uint8
	WK uint64
	WQ uint64
	WR uint64
	WB uint64
	WN uint64
	WP uint64
	BK uint64
	BQ uint64
	BR uint64
	BB uint64
	BN uint64
	BP uint64
	WKmoved uint8
	BKmoved uint8
	WLRmoved uint8
	WRRmoved uint8
	BLRmoved uint8
	BRRmoved uint8
}

type Move struct {
	Name uint8
	Piece uint64
	Start uint64
	End uint64
	Castle uint8
}

func NewBoard() *State {
	var White uint8 = 1
	var WK uint64 = 1<<3
	var WQ uint64 = 1<<4
	var WR uint64 = 1 | 1<<7
	var WB uint64 = 1<<2 | 1<<5
	var WN uint64 = 1<<1 | 1<<6
	var WP uint64 = 1<<8|1<<9|1<<10|1<<11|1<<12|1<<13|1<<14|1<<15
	var BK uint64 = 1<<59
	var BQ uint64 = 1<<60
	var BR uint64 = 1<<56 | 1<<63
	var BB uint64 = 1<<58 | 1<<61
	var BN uint64 = 1<<57 | 1<<62
	var BP uint64 = 1<<48|1<<49|1<<50|1<<51|1<<52|1<<53|1<<54|1<<55
	var No uint8 = 0
	return &State{White, WK, WQ, WR, WB, WN, WP, BK, BQ, BR, BB, BN, BP, No, No, No, No, No, No}
}

func NewMove(name uint8, piece uint64, start uint64, end uint64) *Move {
	return &Move{name, piece, start, end}
}

func (st *State) GenMovesKnight() {
	var knights, team uint64
	if st.White == 1 {
		knights = st.WN
		team = st.AllWhitePieces()
	} else {
		knights = st.BN
		team = st.AllBlackPieces()
	}
	bins := GetPositionsFromBoard(knights)
	moves := []uint64{}
	for _, b := range bins {
		moves = append(moves, KnightMoves(b, team)...)
	}

}

func (st *State) AllWhitePieces() uint64 {
	return st.WK|st.WQ|st.WR|st.WB|st.WN|st.WP
}

func (st *State) AllBlackPieces() uint64 {
	return st.BK|st.BQ|st.BR|st.BB|st.BN|st.BP
}

func (st *State) StateFromMove(mv *Move) *State {
	if st.White == 1 {
		return st.StateFromMoveWhite(mv)
	} else {
		return st.StateFromMoveBlack(mv)
	}
}

func (st *State) StateFromMoveWhite(mv *Move) *State {
	np := mv.Piece&(^mv.Start)
	np |= mv.End
	ns := &State{}

	copier.Copy(ns, st)
	ns.White ^= 1
	if mv.Name == 0 {
		ns.WK = np
		ns.WKmoved = 1
	} else if mv.Name == 1 {
		ns.WQ = np
	} else if mv.Name == 2 {
		ns.WR = np
		if mv.Start == 1 {
			ns.WLRmoved = 1
		}
		if mv.Start == 1<<7 {
			ns.WRRmoved = 1
		}
	} else if mv.Name == 3 {
		ns.WB = np
	} else if mv.Name == 4 {
		ns.WN = np
	} else {
		ns.WP = np
	}
	ns.BK &= ^mv.End
	ns.BQ &= ^mv.End
	ns.BR &= ^mv.End
	ns.BB &= ^mv.End
	ns.BN &= ^mv.End
	ns.BP &= ^mv.End
	return ns
}


func (st *State) StateFromMoveBlack(mv *Move) *State {
	np := mv.Piece&(^mv.Start)
	np |= mv.End
	ns := &State{}

	copier.Copy(ns, st)
	ns.White ^= 1
	if mv.Name == 0 {
		ns.BK = np
		ns.BKmoved = 1
	} else if mv.Name == 1 {
		ns.BQ = np
	} else if mv.Name == 2 {
		ns.BR = np
		if mv.Start == 1<<56 {
			ns.BLRmoved = 1
		}
		if mv.Start == 1<<63 {
			ns.BRRmoved = 1
		}
	} else if mv.Name == 3 {
		ns.BB = np
	} else if mv.Name == 4 {
		ns.BN = np
	} else {
		ns.BP = np
	}
	ns.WK &= ^mv.End
	ns.WQ &= ^mv.End
	ns.WR &= ^mv.End
	ns.WB &= ^mv.End
	ns.WN &= ^mv.End
	ns.WP &= ^mv.End
	return ns
}

func (st *State) GetAllMoves() []*Move {
	if st.White == 1 {
		return st.GetAllMovesWhite()
	} else {
		return st.GetAllMovesBlack()
	}
}

func (st *State) GetAllMovesWhite() []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.WK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, white)
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
		moves := pieces.KnightMoves(n, white)
		for _, m := range moves {
			res = append(res, NewMove(4, st.WN, n, m))
		}
	}
	//Pawns
	pawns := pieces.GetPositionsFromBoard(st.WP)
	for _, p := range pawns {
		moves := pieces.PawnMoves(p, white, black, 1)
		for _, m := range moves {
			res = append(res, NewMove(5, st.WP, p, m))
		}
	}

	return res
}

func (st *State) GetAllMovesBlack() []*Move {
	res := []*Move{}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	//King
	kings := pieces.GetPositionsFromBoard(st.BK)
	for _, k := range kings {
		moves := pieces.KingMoves(k, black)
		for _, m := range moves {
			res = append(res, NewMove(0, st.BK, k, m))
			}
		}
	//Queens
	queens := pieces.GetPositionsFromBoard(st.BQ)
	for _, q := range queens {
		moves := pieces.QueenMoves(q, black, white)
		for _, m := range moves {
			res = append(res, NewMove(1, st.BQ, q, m))
		}
	}
	//Rooks
	rooks := pieces.GetPositionsFromBoard(st.BR)
	for _, r := range rooks {
		moves := pieces.RookMoves(r, black, white)
		for _, m := range moves {
			res = append(res, NewMove(2, st.BR, r, m))
		}
	}
	//Bishops
	bishops := pieces.GetPositionsFromBoard(st.BB)
	for _, b := range bishops {
		moves := pieces.BishopMoves(b, black, white)
		for _, m := range moves {
			res = append(res, NewMove(3, st.BB, b, m))
		}
	}
	//Knights
	knights := pieces.GetPositionsFromBoard(st.BN)
	for _, n := range knights {
		moves := pieces.KnightMoves(n, black)
		for _, m := range moves {
			res = append(res, NewMove(4, st.BN, n, m))
		}
	}
	//Pawns
	pawns := pieces.GetPositionsFromBoard(st.BP)
	for _, p := range pawns {
		moves := pieces.PawnMoves(p, black, white, 0)
		for _, m := range moves {
			res = append(res, NewMove(5, st.BP, p, m))
		}
	}

	return res
}


func GetPositionsFromBoard(piece uint64) []uint64 {
	res := []uint64{}
	var curr, np uint64
	for piece != 0 {
		np = piece&(piece-1)
		curr = np^piece
		fmt.Println(strconv.FormatInt(int64(curr), 2))
		res = append(res, curr)
		piece = np
	}
	return res
}

func (st *State) CanCastleKingWhite bool {
	if st.WKmoved == 1 || !st.WLRmoved == 1 {
		return false
	}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	both := white|black
	need := 1<<1|1<<2
	if need&both != 0 || 1&black != 0 {
		return false
	}
	return true
}

func (st *State) CanCastleQueenWhite bool {
	if st.WKmoved == 1 || !st.WRRmoved == 1 {
		return false
	}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	both := white|black
	need := 1<<4|1<<5|1<<6
	if need&both != 0 || 1<<7&black != 0 {
		return false
	}
	return true
}

func (st *State) CanCastleKingBlack bool {
	if st.BKmoved == 1 || !st.BLRmoved == 1 {
		return false
	}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	both := white|black
	need := 1<<57|1<<58
	if need&both != 0 || 1<<56&white != 0 {
		return false
	}
	return true
}

func (st *State) CanCastleQueenBlack bool {
	if st.BKmoved == 1 || !st.BRRmoved == 1 {
		return false
	}
	white := st.AllWhitePieces()
	black := st.AllBlackPieces()
	both := white|black
	need := 1<<60|1<<61|1<<62
	if need&both != 0 || 1<<63&white != 0 {
		return false
	}
	return true
}

func PawnMoves(bin uint64, same uint64, opp uint64, color uint8) {
	res := []uint64{}
	row, col := GetRowCol(bin)
	both := same|opp
	var curr uint64

	if color == 1 {
		curr = GetBinPos(row+1, col)
		if curr&both == 0 {
			res = append(res, curr)
			if row == 1 {
				curr = GetBinPos(row+2, col)
				if curr&both == 0 {
					res = append(res, curr)
				}
			}
		}
		if col > 0 {
			curr = GetBinPos(row+1, col-1)
			if curr&opp != 0 {
				res = append(res, curr)
			}
		}
		if col < 7 {
			curr = GetBinPos(row+1, col+1)
			if curr&opp != 0 {
				res = append(res, curr)
			}
		}
	} else {
		curr = GetBinPos(row-1, col)
		if curr&both == 0 {
			res = append(res, curr)
			if row == 6 {
				curr = GetBinPos(row-2, col)
				if curr&both == 0 {
					res = append(res, curr)
				}
			}
		}
		if col > 0 {
			curr = GetBinPos(row-1, col-1)
			if curr&opp != 0 {
				res = append(res, curr)
			}
		}
		if col < 7 {
			curr = GetBinPos(row-1, col+1)
			if curr&opp != 0 {
				res = append(res, curr)
			}
		}
	}



}

func KingMoves(bin uint64, same uint64) []uint64 {
	res := []uint64{}
	row, col := GetRowCol(bin)
	var curr uint64
	curr = GetBinPos(row+1, col+1)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row-1, col-1)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row+1, col-1)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row-1, col+1)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row+1, col)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row, col+1)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row-1, col)
	if curr&same == 0{
		res = append(res, curr)
	}
	curr = GetBinPos(row, col-1)
	if curr&same == 0{
		res = append(res, curr)
	}
	return res
}

func QueenMoves(bin uint64, same uint64, opp uint64) []uint64 {
	res := RookMoves(bin, same, opp)
	res = append(res, BishopMoves(bin, same, opp)...)
	return res
}

func RookMoves(bin uint64, same uint64, opp uint64) []uint64 {
	res := []uint64{}
	row, col := GetRowCol(bin)
	var curr, nr, nc uint64
	nr = row+1
	for nr < 8 {
		curr = GetBinPos(nr, col)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nr += 1
	}
	nr = row-1
	for nr >= 0 {
		curr = GetBinPos(nr, col)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nr -= 1
	}
	nc = col+1
	for nc < 8 {
		curr = GetBinPos(row, nc)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nc += 1
	}
	nc = col-1
	for nc >= 0 {
		curr = GetBinPos(row, nc)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nc -= 1
	}
	return res
}

func BishopMoves(bin uint64, same uint64, opp uint64) []uint64 {
	res := []uint64{}
	row, col := GetRowCol(bin)
	var curr, nr, nc uint64
	nr = row+1
	nc = col+1
	for nr < 8 && nc < 8 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nr += 1
		nc += 1
	}
	nr = row-1
	nc = col-1
	for nr >= 0 && nc >= 0 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nr += 1
		nc += 1
	}
	nr = row+1
	nc = col-1
	for nr < 8 && nc >= 0 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nr += 1
		nc -= 1
	}
	nr = row-1
	nc = col+1
	for nr >= 0 && nc < 8 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			res = append(res, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			res = append(res, curr)
		}
		nr += 1
		nc += 1
	}
	return res

}
 
func KnightMoves(bin uint64, team uint64) []uint64 {
	res := []uint64{}
	row, col := GetRowCol(bin)
	var curr, nr, nc uint64
	nr = row+2
	if nr < 8 {
		if col-1 >= 0 {
			curr = GetBinPos(nr, col-1)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
		if col+1 < 8 {
			curr = GetBinPos(nr, col+1)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
	}
	nr = row-2
	if nr < 8 {
		if col-1 >= 0 {
			curr = GetBinPos(nr, col-1)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
		if col+1 < 8 {
			curr = GetBinPos(nr, col+1)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
	}
	nc = col+2
	if nc < 8 {
		if row-1 >= 0 {
			curr = GetBinPos(row-1, nc)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
		if row+1 < 8 {
			curr = GetBinPos(row+1, nc)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
	}
	nc = col-2
	if nc >= 0 {
		if row-1 >= 0 {
			curr = GetBinPos(row-1, nc)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
		if row+1 < 8 {
			curr = GetBinPos(row+1, nc)
			if curr&team == 0 {
				res = append(res, curr)
			}
		}
	}
	return res
}

func GetRowCol(bin uint64) (uint64, uint64) {
	pos := posTable[bin]
	row := pos/8
	col := pos%8
	return row, col
}

func GetBinPos(row uint64, col uint64) uint64 {
	return uint64(math.Pow(2, float64(8*row+col)))
}



func PosTable() map[uint64]uint64 {
	table := make(map[uint64]uint64)
	var curr uint64
	for i := 0; i<64; i++ {
		curr = uint64(math.Pow(2, float64(i)))
		table[curr] = uint64(i)
	}
	return table
}