package pieces

import (
	"math"
)


var (
	posTable = PosTable()
)

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

func GetPositionsFromBoard(piece uint64) []uint64 {
	res := []uint64{}
	var curr, np uint64
	for piece != 0 {
		np = piece&(piece-1)
		curr = np^piece
		res = append(res, curr)
		piece = np
	}
	return res
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