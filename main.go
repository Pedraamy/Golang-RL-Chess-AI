package main

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/algo"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
	"github.com/Pedraamy/Golang-RL-Chess-AI/eval"
	"github.com/Pedraamy/Golang-RL-Chess-AI/misch"
	//"github.com/Pedraamy/Golang-RL-Chess-AI/utils"
	"fmt"
	"strconv"


)


func main() {

	

	/* board := state.NewBoard()
	white := board.AllWhitePieces()
	black := board.AllBlackPieces()
	pieces.PawnMoves(1<<12, white, black, 1)
	i := 0
	var bm *state.Move
	for math.Abs(float64(eval.Eval(board))) < 100000 {
		if board.WK == 0 {
			fmt.Println("Black wins")
			break
		}
		if board.BK == 0 {
			fmt.Println("White wins")
			break
		}
		if i%2 == 0 {
			bm = algo.BestMove(board, 6)
		} else {
			bm = algo.RandomMove(board)
		}
		fmt.Println(Translate(bm.Start, bm.End, bm.Castle))
		fmt.Println(eval.Eval(board))
		board = board.StateFromMove(bm)
		i++ */
	//PlayComp()
	PlayCompAsWhite()

}

func Translate(start uint64, end uint64, castle uint8) string {
	if castle == 1 {
		return "O-O"
	}
	if castle == 2 {
		return "O-O-O"
	}
	var s, e string

	row, col := pieces.GetRowCol(start)

	if col == 0 {
		s += "h"	
	} else if col == 1 {
		s += "g"
	} else if col == 2 {
		s += "f"
	} else if col == 3 {
		s += "e"
	} else if col == 4 {
		s += "d"
	} else if col == 5 {
		s += "c"
	} else if col == 6 {
		s += "b"
	} else if col == 7 {
		s += "a"
	}

	s += strconv.Itoa(row+1)

	row, col = pieces.GetRowCol(end)

	if col == 0 {
		e += "h"	
	} else if col == 1 {
		e += "g"
	} else if col == 2 {
		e += "f"
	} else if col == 3 {
		e += "e"
	} else if col == 4 {
		e += "d"
	} else if col == 5 {
		e += "c"
	} else if col == 6 {
		e += "b"
	} else if col == 7 {
		e += "a"
	}

	e += strconv.Itoa(row+1)

	return s + "  " + e

}

func PlayCompAsWhite() {
	board := state.NewBoard()
	var bm *state.Move
	i := 0
	for {
		if board.WK == 0 {
			fmt.Println("Black wins")
			break
		}
		if board.BK == 0 {
			fmt.Println("White wins")
			break
		}
		if board.White == 1 {
			bm = InputMove(board, 1)
		} else {
			bm = algo.BestMove(board, 6)
		}
		i++
		fmt.Println(Translate(bm.Start, bm.End, bm.Castle))
		fmt.Println(eval.Eval(board))
		board = board.StateFromMove(bm)
	}
	return
}

func PlayCompAsBlack() {
	board := state.NewBoard()
	var bm *state.Move
	i := 0
	for {
		if board.WK == 0 {
			fmt.Println("Black wins")
			break
		}
		if board.BK == 0 {
			fmt.Println("White wins")
			break
		}
		if board.White == 1 {
			bm = algo.BestMove(board, 6)
		} else {
			bm = InputMove(board, 0)
		}
		i++
		fmt.Println(Translate(bm.Start, bm.End, bm.Castle))
		fmt.Println(eval.Eval(board))
		board = board.StateFromMove(bm)
	}
	return
}

func InputMove (st *state.State, color uint8) *state.Move {
	var castle string
    fmt.Println("Castle?")
	fmt.Scan(&castle)
	for castle != "Q" && castle != "K" && castle != "no" {
		fmt.Scanln(&castle)
    	fmt.Println("Castle?")
	}
	if castle == "Q" {
		return state.NewMove(1, 1, 1, 1, 2)
	} else if castle == "K" {
		return state.NewMove(1, 1, 1, 1, 1)
	}
	var name uint8
	var piece, start, end uint64
    fmt.Println("Name?")
	fmt.Scan(&name)

    fmt.Println("Start?")
	fmt.Scan(&start)
	start = 1<<start

    fmt.Println("End?")
	fmt.Scan(&end)
	end = 1<<end

	if name == 0 {
		if color == 0 {
			piece = st.BK
		} else {
			piece = st.WK
		}

	} else if name == 1 {
		if color == 0 {
			piece = st.BQ
		} else {
			piece = st.WQ
		}

	} else if name == 2 {
		if color == 0 {
			piece = st.BR
		} else {
			piece = st.WR
		}

	} else if name == 3 {
		if color == 0 {
			piece = st.BB
		} else {
			piece = st.WB
		}

	} else if name == 4 {
		if color == 0 {
			piece = st.BN
		} else {
			piece = st.WN
		}

	} else {
		if color == 0 {
			piece = st.BP
		} else {
			piece = st.WP
		}
	}
	return state.NewMove(name, piece, start, end, 0)

}


