package main

import (
	"github.com/Pedraamy/Golang-RL-Chess-AI/state"
	"github.com/Pedraamy/Golang-RL-Chess-AI/algo"
	"github.com/Pedraamy/Golang-RL-Chess-AI/pieces"
	"fmt"
	"strconv"
	"math"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

)


func main() {
	PosTable := pieces.PosTable()
	

	board := state.NewBoard()
	white := board.AllWhitePieces()
	black := board.AllBlackPieces()
	pieces.PawnMoves(1<<12, white, black, 1)
	i := 0
	var bm *state.Move
	for math.Abs(algo.Eval(board)) < 5000 {
		if i%2 == 0 {
			bm = algo.BestMove(board, 4)
		} else {
			bm = algo.RandomMove(board)
		}
		fmt.Println(Translate(bm.Name, bm.End, bm.Castle))
		fmt.Println(algo.Eval(board))
		board = board.StateFromMove(bm)
		i++
	}
	mv := algo.BestMove(board, 4)
	fmt.Println(mv.Name)
	fmt.Println(PosTable[mv.Start], PosTable[mv.End])
	
}

func Translate(name uint8, end uint64, castle uint8) string {
	if castle == 1 {
		return "O-O"
	}
	if castle == 2 {
		return "O-O-O"
	}
	var res string
	if name == 0 {
		res += "K"
	} else if name == 1 {
		res += "Q"
	} else if name == 2 {
		res += "R"
	} else if name == 3 {
		res += "B"
	} else if name == 4 {
		res += "N"
	}

	row, col := pieces.GetRowCol(end)

	if col == 0 {
		res += "h"	
	} else if col == 1 {
		res += "g"
	} else if col == 2 {
		res += "f"
	} else if col == 3 {
		res += "e"
	} else if col == 4 {
		res += "d"
	} else if col == 5 {
		res += "c"
	} else if col == 6 {
		res += "b"
	} else if col == 7 {
		res += "a"
	}

	res += strconv.Itoa(row+1)

	return res

}

func ChessBoard() {
	const chessboardSquaresPerSide = 8
	const chessboardSquares = chessboardSquaresPerSide * chessboardSquaresPerSide

	const chessboardSquareSideSize = 120
	const chessboardSideSize = chessboardSquaresPerSide * chessboardSquareSideSize

	chessboardRect := image.Rect(0, 0, chessboardSideSize, chessboardSideSize)
	chessboardImage := image.NewRGBA(chessboardRect)
	uniform := &image.Uniform{}

	for s := 0; s < chessboardSquares; s++ {
		x := s / chessboardSquaresPerSide
		y := s % chessboardSquaresPerSide

		x0 := x * chessboardSquareSideSize
		y0 := y * chessboardSquareSideSize
		x1 := x0 + chessboardSquareSideSize
		y1 := y0 + chessboardSquareSideSize

		rect := image.Rect(x0, y0, x1, y1)

		if isBlack := (x+y)%2 == 1; isBlack {
			uniform.C = color.Black
		} else {
			uniform.C = color.White
		}

		draw.Draw(chessboardImage, rect, uniform, image.Point{}, draw.Src)
	}

	chessboardFile, err := os.Create("chessboard.png")
	if err != nil {
		panic(err)
	}

	if err := png.Encode(chessboardFile, chessboardImage); err != nil {
		chessboardFile.Close()
		panic(err)
	}

	if err := chessboardFile.Close(); err != nil {
		panic(err)
	}
}
