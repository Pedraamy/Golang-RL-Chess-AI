package game

type Pos struct {
	x uint8
	y uint8
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
	
}