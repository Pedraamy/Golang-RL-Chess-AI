package game

import (
	"math"
)

var (
	posTable = PosTable()
)

type Pos struct {
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


type Game struct {
	WhitePieces map[string]Pos
	BlackPieces map[string]Pos	
	Board [8][8]uint8
	WKmoved bool
	BKmoved bool
	WLRmoved bool
	WRRMoved bool
	BLRmoved bool
	BRRmoved bool
}

func (game *Game) GetMoves() []int {
	return []int{}
}

func (game *Game) GenMovesKing() {

}

func GetPosition(piece uint64) []int {
	res := []int{}
	if piece == 0 {
		return res
	}
	var curr uint64
	for piece&(piece-1) != 0 {
		curr = piece&(piece-1)
		res = append(res, posTable[curr])
		piece ^= curr
	}
	res = append(res, posTable[piece])
	return res


}

func PosTable() map[uint64]int {
	table := make(map[uint64]int)
	var curr uint64
	for i := 0; i<64; i++ {
		curr = uint64(math.Pow(2, float64(i)))
		table[curr] = i
	}
	return table
}