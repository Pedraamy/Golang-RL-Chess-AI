package move

type Move struct {
	Name uint8
	Piece uint64
	Start uint64
	End uint64
}

func NewMove(name uint8, piece uint64, start uint64, end uint64) *Move {
	return &Move{name, piece, start, end}
}