package pieces

import (
	"math"
	//"fmt"
)


var (
	PosTable = posTable()
)

func PawnMoves(bin uint64, same uint64, opp uint64, color uint8) ([]uint64, []uint64) {
	moves := []uint64{}
	captures := []uint64{}
	row, col := GetRowCol(bin)
	both := same|opp
	var curr uint64

	if color == 1 {
		if row < 6 {
			curr = GetBinPos(row+1, col)
			if curr&both == 0 {
				moves = append(moves, curr)
				if row == 6 {
	
				}
				if row == 1 {
					curr = GetBinPos(row+2, col)
					if curr&both == 0 {
						moves = append(moves, curr)
					}
				}
			}
			if col > 0 {
				curr = GetBinPos(row+1, col-1)
				if curr&opp != 0 {
					captures = append(captures, curr)
				}
			}
			if col < 7 {
				curr = GetBinPos(row+1, col+1)
				if curr&opp != 0 {
					captures = append(captures, curr)
				}
			}	
		}
	} else {
		if row > 1 {
			curr = GetBinPos(row-1, col)
			if curr&both == 0 {
				moves = append(moves, curr)
				if row == 6 {
					curr = GetBinPos(row-2, col)
					if curr&both == 0 {
						moves = append(moves, curr)
					}
				}
			}
			if col > 0 {
				curr = GetBinPos(row-1, col-1)
				if curr&opp != 0 {
					captures = append(captures, curr)
				}
			}
			if col < 7 {
				curr = GetBinPos(row-1, col+1)
				if curr&opp != 0 {
					captures = append(captures, curr)
				}
			}
		}
	}
	return captures, moves
}

func KingMoves(bin uint64, same uint64, opp uint64) ([]uint64, []uint64) {
	moves := []uint64{}
	captures := []uint64{}
	row, col := GetRowCol(bin)
	var curr uint64
	if row+1 < 8 {
		curr = GetBinPos(row+1, col)
		if curr&opp != 0 {
			captures = append(captures, curr)
		} else if curr&same == 0{
			moves = append(moves, curr)
		}
		if col+1 < 8 {
			curr = GetBinPos(row+1, col+1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0{
				moves = append(moves, curr)
			}
		}
		if col-1 >= 0 {
			curr = GetBinPos(row+1, col-1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0{
				moves = append(moves, curr)
			}
		}
	}
	if row-1 >= 0 {
		curr = GetBinPos(row-1, col)
		if curr&opp != 0 {
			captures = append(captures, curr)
		} else if curr&same == 0{
			moves = append(moves, curr)
		}
		if col+1 < 8 {
			curr = GetBinPos(row-1, col+1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0{
				moves = append(moves, curr)
			}
		}
		if col-1 >= 0 {
			curr = GetBinPos(row-1, col-1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0{
				moves = append(moves, curr)
			}
		}
	}
	if col+1 < 8 {
		curr = GetBinPos(row, col+1)
		if curr&opp != 0 {
			captures = append(captures, curr)
		} else if curr&same == 0{
			moves = append(moves, curr)
		}
	}
	if col-1 >= 0 {
		curr = GetBinPos(row, col-1)
		if curr&opp != 0 {
			captures = append(captures, curr)
		} else if curr&same == 0{
			moves = append(moves, curr)
		}
	}
	return captures, moves
}

func QueenMoves(bin uint64, same uint64, opp uint64) ([]uint64, []uint64) {
	captures, moves := RookMoves(bin, same, opp)
	cap2, mov2 := BishopMoves(bin, same, opp)
	captures = append(captures, cap2...)
	moves = append(moves, mov2...)
	return captures, moves
}

func RookMoves(bin uint64, same uint64, opp uint64) ([]uint64, []uint64) {
	moves := []uint64{}
	captures := []uint64{}
	row, col := GetRowCol(bin)
	var curr uint64
	var nr, nc int
	nr = row+1
	for nr < 8 {
		curr = GetBinPos(nr, col)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nr += 1
	}
	nr = row-1
	for nr >= 0 {
		curr = GetBinPos(nr, col)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nr -= 1
	}
	nc = col+1
	for nc < 8 {
		curr = GetBinPos(row, nc)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nc += 1
	}
	nc = col-1
	for nc >= 0 {
		curr = GetBinPos(row, nc)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nc -= 1
	}
	return captures, moves
}

func BishopMoves(bin uint64, same uint64, opp uint64) ([]uint64, []uint64) {
	moves := []uint64{}
	captures := []uint64{}
	row, col := GetRowCol(bin)
	var curr uint64
	var nr, nc int
	nr = row+1
	nc = col+1
	for nr < 8 && nc < 8 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nr += 1
		nc += 1
	}
	nr = row-1
	nc = col-1
	for nr >= 0 && nc >= 0 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nr -= 1
		nc -= 1
	}
	nr = row+1
	nc = col-1
	for nr < 8 && nc >= 0 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nr += 1
		nc -= 1
	}
	nr = row-1
	nc = col+1
	for nr >= 0 && nc < 8 {
		curr = GetBinPos(nr, nc)
		if curr&opp != 0{
			captures = append(captures, curr)
			break
		} else if curr&same != 0{
			break
		} else {
			moves = append(moves, curr)
		}
		nr -= 1
		nc += 1
	}
	return captures, moves

}
 
func KnightMoves(bin uint64, same uint64, opp uint64) ([]uint64, []uint64) {
	moves := []uint64{}
	captures := []uint64{}
	row, col := GetRowCol(bin)
	var curr uint64
	var nr, nc int
	nr = row+2
	if nr < 8 {
		if col-1 >= 0 {
			curr = GetBinPos(nr, col-1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
		if col+1 < 8 {
			curr = GetBinPos(nr, col+1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
	}
	nr = row-2
	if nr >= 0 {
		if col-1 >= 0 {
			curr = GetBinPos(nr, col-1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
		if col+1 < 8 {
			curr = GetBinPos(nr, col+1)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
	}
	nc = col+2
	if nc < 8 {
		if row-1 >= 0 {
			curr = GetBinPos(row-1, nc)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
		if row+1 < 8 {
			curr = GetBinPos(row+1, nc)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
	}
	nc = col-2
	if nc >= 0 {
		if row-1 >= 0 {
			curr = GetBinPos(row-1, nc)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
		if row+1 < 8 {
			curr = GetBinPos(row+1, nc)
			if curr&opp != 0 {
				captures = append(captures, curr)
			} else if curr&same == 0 {
				moves = append(moves, curr)
			}
		}
	}
	return captures, moves
}

func GetPiecesFromBoard(piece uint64) []uint64 {
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

func GetRowCol(bin uint64) (int, int) {
	pos := PosTable[bin]
	row := int(pos/8)
	col := int(pos%8)
	return row, col
}

func GetBinPos(row int, col int) uint64 {
	return uint64(math.Pow(2, float64(8*row+col)))
}

func posTable() map[uint64]uint64 {
	table := make(map[uint64]uint64)
	for i := 0; i<64; i++ {
		table[1<<i] = uint64(i)
	}
	return table
}